package hendler

import (
	"my_Ipa/generated/bus"
	"my_Ipa/generated/wheather"
)

type Handler struct {
	Transport bus.TransportServiceClient
	Weather   wheather.WitherServerClient
}

func NewHendler(weather wheather.WitherServerClient, transport bus.TransportServiceClient) *Handler {
	return &Handler{
		Transport: transport,
		Weather:   weather,
	}
}
