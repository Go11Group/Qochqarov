package hendler

import (
	"fmt"
	pb "my_Ipa/generated/wheather"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CurrentWeather(c *gin.Context) {
	location := c.Query("location")

	resp, err := h.Weather.GetCurrentWeather(c, &pb.CurrentWeatherRequest{Location: location})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) WeatherForecast(c *gin.Context) {
	location := c.Query("location")
	day := c.Query("day")
	var day2 int
	if day != "" {
		k, err := strconv.Atoi(day)
		day2 = k
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.Weather.GetWeatherForecast(c, &pb.WeatherForecastRequest{Location: location, Day: int32(day2)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetReportWeatherCondition(c *gin.Context) {
	location := c.Query("location")
	date := c.Query("date")
	fmt.Println(location,date)
	resp, err := h.Weather.ReportWeatherCondition(c, &pb.IsTrafficRequest{Location: location, Date: date})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
