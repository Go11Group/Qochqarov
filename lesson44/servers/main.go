package main

import (
	"context"
	"log"
	"net"

	pb "my_mod/translater"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTranslaterServer
	Hashmap map[string]string
}

func main() {
	translations := map[string]string{
		"apple":      "olma",
		"banana":     "banan",
		"cherry":     "gilos",
		"date":       "xurmo",
		"elderberry": "qarag'at",
		"fig":        "anjir",
		"grape":      "uzum",
		"honeydew":   "qovun",
		"kiwi":       "kiwi",
		"lemon":      "limon",
		"mango":      "mango",
		"nectarine":  "nektarin",
		"orange":     "apelsin",
		"papaya":     "papaya",
		"quince":     "behi",
		"raspberry":  "malina",
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTranslaterServer(s, &Server{Hashmap: translations})

	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
}

func (s *Server) GetTranslaterText(ctx context.Context, req *pb.RequeTranslater) (*pb.ResponceTrablater, error) {
	response := make(map[string]string)
	for _, v := range req.StrUzb {
		if translation, ok := s.Hashmap[v]; ok {
			response[v] = translation
		} else {
			response[v] = "topilmadi"
		}
	}
	return &pb.ResponceTrablater{StrEng: response}, nil
}
