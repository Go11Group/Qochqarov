package main

import (
	"context"
	"fmt"
	pb "my_mod/translater"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewTranslaterClient(conn)
	conte, consel := context.WithTimeout(context.Background(), time.Second)
	defer consel()

	slc := []string{"apple", "banana", "gilos"}

	req := &pb.RequeTranslater{
		StrUzb: slc,
	}

	res, err := c.GetTranslaterText(conte, req)
	if err != nil {
		panic(err)
	}

	for k, v := range res.StrEng {
		fmt.Printf("%s: %s\n", k, v)
	}
}
