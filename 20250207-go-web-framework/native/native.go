package native

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request")
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	select {
	case <-r.Context().Done():
		fmt.Println("context done")
	case <-time.After(10 * time.Second):
		fmt.Println("timeout")
	}
}

func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler) // Correct route pattern
	logging := AuthMiddleware(mux)
	srv := &http.Server{Addr: ":8080", Handler: logging}
	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()
		<-ctx.Done()
		srv.Shutdown(context.Background())
	}()
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
