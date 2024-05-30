package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterUser(c echo.Context) error {
	UserRegistered.Inc()
	return c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

func ReadArticle(c echo.Context) error {
	userId := c.Request().Header.Get("Prom-Demo-User-Id")
	if userId != "" {
		UserReadArticle.With(map[string]string{"user_id": userId}).Inc()
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func PromDurationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		respStatus := c.Response().Status
		duration := time.Since(start)
		RequestDuration.With(map[string]string{
			"response_status": strconv.Itoa(respStatus),
		}).Observe(duration.Seconds())
		RequestTotal.With(map[string]string{
			"response_status": strconv.Itoa(respStatus),
		}).Inc()
		return err
	}
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	promServer := &http.Server{
		Addr: ":8081",
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(PromDurationMiddleware)
	e.POST("/user/register", RegisterUser)
	e.GET("/article", ReadArticle)

	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	go func() {
		<-sig
		slog.Info("shutting down application")
		cancel()
	}()

	// start prometheus exporter
	go func() {
		slog.Info("Listening and serving prometheus exporter", slog.Int("port", 8081))
		if err := promServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed listening prometheus exporter", slog.Any("err", err))
			panic(err)
		}
	}()

	// start echo server
	go func() {
		slog.Info("Listening and serving HTTP", slog.Int("port", 8080))
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed starting HTTP server", slog.Any("err", err))
			panic(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		slog.Error("HTTP Server shutdown failed", slog.Any("err", err))
		panic(err)
	}
	if err := promServer.Shutdown(shutdownCtx); err != nil {
		slog.Error("Prometheus shutdown failed", slog.Any("err", err))
		panic(err)
	}
	slog.Info("application shut down")
}
