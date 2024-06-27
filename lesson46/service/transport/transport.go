package transport

import (
	pb "my_mod/generated/wheather"
	"my_mod/storege/postgres"
)

type Server struct {
	pb.UnimplementedWitherServerServer
	buss postgres.Weatherrepo
}

func (s *Server) GetBusSchedule() {

}
