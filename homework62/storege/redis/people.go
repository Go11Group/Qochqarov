package redis

import (
	"context"
	"encoding/json"
	"log"
	pb "my_mod/generated/users"
	"time"

	"github.com/redis/go-redis/v9"
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

	repo.RD.Set(context.Background(), in.Id, maps, time.Duration(t.Add(30*time.Minute).Unix()))

	return &pb.CresteUserResponse{}, nil
}

func (repo *UserRepo) GetByIdUser(in *pb.GetByIdUserRequest) (*pb.GetByIdUserResponse, error) {
	val, err := repo.RD.Get(context.Background(), in.UserId).Result()
	if err != nil {
		log.Fatalf("Redisdan ma'lumot olishda xatolik: %v", err)
	}

	var user pb.GetByIdUserResponse
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (repo *UserRepo) DeketeUser(id string) (*pb.DeleteUserResponse, error) {
	err := repo.RD.Del(context.Background(), id).Err()
	return &pb.DeleteUserResponse{Success: true}, err
}

func (repo *UserRepo) UpdateUser(in *pb.UpdateUserrequest) (*pb.UpdateUserResponse, error) {

	user, err := json.Marshal(in)
	if err != nil {
		return &pb.UpdateUserResponse{Success: false}, err
	}
	err = repo.RD.Set(context.Background(), in.UserId, user, 0).Err()
	return &pb.UpdateUserResponse{Success: true}, err
}
