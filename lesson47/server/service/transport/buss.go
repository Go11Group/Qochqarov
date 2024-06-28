package transport

import (
	"context"
	pb "my_mod/generated/bus"
	postgres "my_mod/storege"
)

type Servers struct {
	pb.UnimplementedTransportServiceServer
	Buss *postgres.Bussrepo
}

func (s *Servers) GetBusSchedule(ctx context.Context, in *pb.BusRequest) (*pb.BusScheduleResponse, error) {

	resp, err := s.Buss.BusSchedule(in)

	return resp, err

}
func (s *Servers) TrackBusLocation(ctx context.Context, in *pb.BusRequest) (*pb.BusLocationResponse, error) {

	resp, err := s.Buss.TraBusLocations(in)

	return resp, err
}
func (s *Servers) ReportTrafficJam(ctx context.Context, in *pb.TrafficJamReportRequest) (*pb.TrafficJamReportResponse, error) {

	resp, err := s.Buss.GetReportTrafficJam(in)

	return resp, err
}
