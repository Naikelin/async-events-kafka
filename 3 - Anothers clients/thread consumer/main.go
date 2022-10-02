package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

// critical section
var s []string
var count int
var mu sync.Mutex

func printSlice(a []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}

func saveInSlice(a string) {
	mu.Lock()
	defer mu.Unlock()
	count++
	s = append(s, a)
}

func readElements(c chan bool) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "CG1",
		Topic:    "register-event",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	// with deadline 30secs
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Fatal("failed to fetch message:", err)
			break
		}

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit message:", err)
			break
		}

		saveInSlice(string(m.Value))
	}

	c <- true

}

func main() {
	c := make(chan bool)
	go readElements(c)

	for {
		time.Sleep(time.Second * 10)
		printSlice(s)
	}

	//<-c
}
