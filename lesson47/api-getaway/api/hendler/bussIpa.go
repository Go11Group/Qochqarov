package hendler

import (
	pb "my_Ipa/generated/bus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBusScheduleIpa(c *gin.Context) {
	number := c.Query("bussNumber")

	resp, err := h.Transport.GetBusSchedule(c,&pb.BusRequest{BusNumber: number})
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler)   GetTrackBusLocation(c *gin.Context){
	number := c.Query("bussNumber")

	resp, err := h.Transport.TrackBusLocation(c,&pb.BusRequest{BusNumber: number})
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


func (h *Handler)   GetReportTrafficJam(c *gin.Context){
	location := c.Query("location")
	discription:=c.Query("description")	

	resp, err := h.Transport.ReportTrafficJam(c,&pb.TrafficJamReportRequest{Locatsion: location,Discription: discription})
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


