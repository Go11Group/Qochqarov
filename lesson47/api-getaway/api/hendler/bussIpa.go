package hendler

import ("github.com/gin-gonic/gin"
	pb "my_Ipa/generated/bus"
)

func (h *Handler) GetBusScheduleIpa(c *gin.Context) {
	number:=c.Param("bussNumber")

	h.Transport.
}