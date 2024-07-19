package redis

import (
	"context"
	pb "my_mod/generated/users"
	"my_mod/storege/redis"
	"time"

	"github.com/redis/go-redis"
)

type UserRepo struct {
	RD *redis.Client
}

func NewUserRepo(rd *redis.Client) *UserRepo {
	return &UserRepo{RD: rd}
}


func (repo *UserRepo) Create(in *pb.CresteUserRequest) (*pb.CresteUserResponse, error) {
	var t time.Time

	maps := map[string]interface{}{
		"name":   in.Name,
		"age":    in.Age,
		"gendor": in.Gendor,
		"year":   in.Year,
	}

	repo.RD.HSet(context.Background(), in.Name, maps)
	repo.RD.Expire(context.Background(), in.Name, time.Duration(t.Add(30*time.Minute).Unix()))

	return &pb.CresteUserResponse{}, nil
}
