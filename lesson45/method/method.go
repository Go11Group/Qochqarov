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

func (s *Server) AddBook(ctx context.Context, r *pb.AddBookRequest) (*pb.AddBookRespons, error) {
	id, err := s.Library.LibaryCreate(r)
	if err != nil {
		return nil, err
	}

	return &pb.AddBookRespons{BookId: *id}, nil
}

func (s *Server) SearchBook(ctx context.Context, r *pb.SearchBookRequest) (*pb.SearchBookRespons, error) {
	book, err := s.Library.LIbaryGet(r.Query)

	if err != nil {
		return nil, err
	}

	return &pb.SearchBookRespons{Books: book}, nil
}

func (s *Server) BorrowBook(ctx context.Context, r *pb.BorrowBookRequest) (*pb.BorrowBookRespons, error) {
	borw,err:=s.Library.LibaryBorrow(r)

	if err != nil {
		return nil, err
	}

	return borw,nil
}
