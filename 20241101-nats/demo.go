package main

import (
	"context"
	"log"
	"os/signal"
	"time"
)

func ChannelSelect() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		signal.Notify()
		// SIGTERM / SIGINT
		cancel()
	}()
	select {
	case <-ctx.Done():
		exit()
	case <-time.After(time.Second):
		doWOrk()
	}
}

func doWOrk() error {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	panic("WOW SAYA PANIK")
}
