package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const topicName = "mytopic"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	switch os.Args[1] {
	case "producer":
		writer := &kafka.Writer{
			Addr:         kafka.TCP("localhost:9092"),
			Async:        true,
			RequiredAcks: kafka.RequireOne,
			BatchSize:    100,
			BatchTimeout: 5 * time.Second,
			Compression:  kafka.Gzip,
		}
		defer writer.Close()
		for {
			fmt.Print("> ")
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			response = strings.TrimSuffix(response, "\n")
			if response == "" {
				break
			}
			if err := writer.WriteMessages(ctx, kafka.Message{
				Topic: topicName,
				Value: []byte(response),
				Key:   nil,
			}); err != nil {
				if errors.Is(err, context.Canceled) {
					break
				}
				panic(err)
			}
		}
	case "consumer":
		groupId := os.Args[2]
		consumer := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   topicName,
			GroupID: groupId,
		})
		for {
			m, err := consumer.FetchMessage(ctx)
			if errors.Is(err, context.Canceled) {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(string(m.Value))
			err = consumer.CommitMessages(ctx, m)
			if err != nil {
				panic(err)
			}
		}
	}
}
