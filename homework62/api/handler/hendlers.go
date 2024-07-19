package handler

import "my_mod/generated/users"

type Handler struct {
	User users.UserServiceClient
}

func NewHandler(user users.UserServiceClient) *Handler {
	return &Handler{
		User: user,
	}
}
