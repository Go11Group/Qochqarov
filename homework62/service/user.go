package service

import (
	"context"
	pb "my_mod/generated/users"
	"my_mod/storege/redis"
)


type UserService struct{
	UserRepo *redis.UserRepo
}

func NewUserRepo(user *redis.UserRepo) *UserService  {
	return &UserService{UserRepo: user}
}

func (u UserService) CresteUser(ctx *context.Context,in *pb.CresteUserRequest) (*pb.CresteUserResponse,error){
	return u.UserRepo.Create(in)
}

func (u UserService) GetByIdUser(ctx *context.Context,in *pb.GetByIdUserRequest) (*pb.GetByIdUserResponse,error){
	return u.UserRepo.GetByIdUser(in)
}

func (u UserService) UpdateUser(ctx *context.Context,in *pb.UpdateUserrequest) (*pb.UpdateUserResponse,error){
	return u.UserRepo.UpdateUser(in)
}

func (u UserService) DeleteUser(ctx *context.Context,in *pb.DeleteUserRequest) (*pb.DeleteUserResponse,error){
	return u.UserRepo.DeketeUser(in.UserId)
}