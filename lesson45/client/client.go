package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "my_mod/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	c := pb.NewLibraryServerClient(conn)

	contex, consel := context.WithTimeout(context.Background(), time.Second)

	defer consel()

	reqAdd := pb.AddBookRequest{
		Titile: "vatan",
		Author: "hamid olimjon",
		Year:   1960,
	}

	resAdd, err := c.AddBook(contex, &reqAdd)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("salom",resAdd)
//---------------------------------------------------
	reqSorch := pb.SearchBookRequest{
		Query: "To Kill a Mockingbird",
	}

	resSorch, err := c.SearchBook(contex, &reqSorch)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("alik",resSorch)
	//-----------------------------------------------

	reqBorrow := pb.BorrowBookRequest{
		BookId: "49eeae83-561a-40d1-820e-7adc8636f3f4",
		UserId: "U001",
	}

	resBorrow, err := c.BorrowBook(contex, &reqBorrow)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("nima",resBorrow)



}
