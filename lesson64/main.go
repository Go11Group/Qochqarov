package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	//create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//create a new pubsub client
	pubsub := rdb.Subscribe(ctx, "Amazon", "Tesla")
	defer pubsub.Close()

	// Wait for confirmation that subscription is created
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	// Start a goroutine to receive messages
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Printf("Time : %s\nCurrent Price : %s\nCompany Name : %s\n", time.Now().Format("2006-01-02 15:04:05"), msg.Payload, msg.Channel)
	}
	
}