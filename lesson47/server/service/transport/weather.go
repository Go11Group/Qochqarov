package transport

import (
	"context"
	pb "my_mod/generated/wheather"
	postgres "my_mod/storege"
)

type Server struct {
	pb.UnimplementedWitherServerServer
	With *postgres.Weatherrepo
}

func (s *Server) GetCurrentWeather(ctx context.Context, in *pb.CurrentWeatherRequest) (*pb.CurrentWeatherRespons, error) {

	resp, err := s.With.GetWeather(in)

	return resp, err

}
func (s *Server) GetWeatherForecast(ctx context.Context, in *pb.WeatherForecastRequest) (*pb.WeatherForecastRespons, error) {

	resp, err := s.With.WeatherForecast(in)

	return resp, err

}
func (s *Server) ReportWeatherCondition(ctx context.Context, in *pb.IsTrafficRequest) (*pb.IsTraficRespons, error) {

	resp, err := s.With.IsTraffic(in)

	return resp, err
}
