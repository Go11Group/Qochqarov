package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Brand struct {
	Name  string
	Price string
	Time  string
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Publish messages to different channels
	for true {
		err := publishRedis(ctx, rdb, "Amazon")
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Messages published successfully")
}

func publishRedis(ctx context.Context, rdb *redis.Client, channel string) error {
	brand := Random(channel)
	data, err := json.Marshal(brand)
	if err != nil {
		panic(err)
	}
	return rdb.Publish(ctx, channel, data).Err()
}

func Random(name string) Brand {
	brand := Brand{
		Name:  name,
		Price: "$" + strconv.Itoa(rand.Intn(1000)),
		Time:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return brand
}