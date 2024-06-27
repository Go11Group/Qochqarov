package transport

import (
	"context"
	pb "my_mod/generated/bus"
	postgres "my_mod/storege"

	"google.golang.org/grpc"
)

type Servers struct {
	pb.UnimplementedTransportServiceServer
	buss *postgres.Bussrepo
}

func (s *Servers) GetBusSchedule(ctx context.Context, in *pb.BusRequest, opts ...grpc.CallOption) (*pb.BusScheduleResponse, error) {

	resp, err := s.buss.BusSchedule(in)

	return resp, err

}
func (s *Servers) TrackBusLocation(ctx context.Context, in *pb.BusRequest, opts ...grpc.CallOption) (*pb.BusLocationResponse, error) {

	resp, err := s.buss.TraBusLocations(in)

	return resp, err
}
func (s *Servers) ReportTrafficJam(ctx context.Context, in *pb.TrafficJamReportRequest, opts ...grpc.CallOption) (*pb.TrafficJamReportResponse, error) {

	resp, err := s.buss.GetReportTrafficJam(in)

	return resp, err
}
