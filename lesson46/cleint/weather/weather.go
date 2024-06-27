package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "my_mod/generated/wheather"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	c := pb.NewWitherServerClient(conn)

	contex, consel := context.WithTimeout(context.Background(), time.Second)

	defer consel()
	reqCurrent := pb.CurrentWeatherRequest{
		Location: "tatu",
	}
	respCurrent, err := c.GetCurrentWeather(contex, &reqCurrent)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(respCurrent)
	//-----------------------------------------

	reqForecas:=pb.WeatherForecastRequest{
			Location: "bodomzor",
			Day: 4,
	}
	respForecas,err:=c.GetWeatherForecast(contex,&reqForecas)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respForecas)	
	//------------------------------------------


	reqReport:=pb.IsTrafficRequest{
		Location: "minor",
		Date: "20-04-2023",
	}

	respReport,err:=c.ReportWeatherCondition(contex,&reqReport)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(respReport)

}
