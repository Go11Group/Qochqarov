package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "my_mod/generated/bus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	c := pb.NewTransportServiceClient(conn)

	contex, consel := context.WithTimeout(context.Background(), time.Second)

	defer consel()

	reqbusnum := pb.BusRequest{
		BusNumber: "044",
	}

	respBusNum, err := c.GetBusSchedule(contex, &reqbusnum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respBusNum)
	//---------------------------------------

	respTrack,err:=c.TrackBusLocation(contex,&reqbusnum)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(respTrack)

	//-------------------------------

	reqTraf:=pb.TrafficJamReportRequest{
		Locatsion: "bodomzor",
		Discription: "kjirjghrekrwmr",
	}

	respTraf,err:=c.ReportTrafficJam(contex,&reqTraf)

	fmt.Println(respTraf)
}
