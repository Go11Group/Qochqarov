package transport

import (
	"context"
	pb "my_mod/generated/wheather"
	postgres "my_mod/storege"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedWitherServerServer
	with *postgres.Weatherrepo
}

func (s *Server) GetCurrentWeather(ctx context.Context, in *pb.CurrentWeatherRequest, opts ...grpc.CallOption) (*pb.CurrentWeatherRespons, error) {

	resp, err := s.with.GetWeather(in)

	return resp, err

}
func (s *Server) GetWeatherForecast(ctx context.Context, in *pb.WeatherForecastRequest, opts ...grpc.CallOption) (*pb.WeatherForecastRespons, error) {

	resp, err := s.with.WeatherForecast(in)

	return resp, err

}
func (s *Server) ReportWeatherCondition(ctx context.Context, in *pb.IsTrafficRequest, opts ...grpc.CallOption) (*pb.IsTraficRespons, error) {
	
	resp,err:=s.with.IsTraffic(in)

	return resp,err
}

