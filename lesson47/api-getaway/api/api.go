package api

import (
	"my_Ipa/api/hendler"
	"my_Ipa/generated/bus"
	"my_Ipa/generated/wheather"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Router(conn *grpc.ClientConn) *gin.Engine {

	weather := wheather.NewWitherServerClient(conn)
	transport := bus.NewTransportServiceClient(conn)
	router := gin.Default()
	handler := hendler.NewHendler(weather, transport)

	router.GET("/weather/get/1", handler.CurrentWeather)
	router.GET("/weather/get/2", handler.WeatherForecast)
	router.GET("/weather/get/3", handler.GetReportWeatherCondition)

	router.GET("/transport/get/1", handler.GetBusScheduleIpa)
	router.GET("/transport/get/2", handler.GetTrackBusLocation)
	router.GET("/transport/get/3", handler.GetReportTrafficJam)



	return router
}
