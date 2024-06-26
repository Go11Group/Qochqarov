package method

import (
	"context"
	pb "my_mod/library"
	"my_mod/postgres"
)

type Server struct {
	pb.UnimplementedLibraryServerServer

	Library *postgres.LibarySorage
}

func (s *Server) AddBook(ctx *context.Context, r *pb.AddBookRequest) (*pb.AddBookRespons, error) {
	id, err := s.Library.LibaryCreate(r)
	if err != nil {
		return nil, err
	}

	return &pb.AddBookRespons{BookId: *id}, context.Background().Err()
}

func (s *Server) SearchBook(ctx *context.Context, r *pb.SearchBookRequest) (*pb.SearchBookRespons, error) {

}

func (s *Server) BorrowBook(ctx *context.Context, r *pb.BorrowBookRequest) (*pb.BorrowBookRespons, error) {

}
