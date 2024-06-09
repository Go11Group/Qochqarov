package handler

import (
	"encoding/json"
	"my_mod/model"
	"net/http"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	var users model.User
	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&users)
		if err != nil {
			panic(err)
		}
		h.Userss.U_Create(users)
	case http.MethodGet:
		user, err := h.Userss.U_Read(users)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			panic(err)
		}
	case http.MethodPut:
		err:=json.NewDecoder(r.Body).Decode(&users)
		if err != nil {
			panic(err)
		}
		h.Userss.U_Update(users.Age,users.First_name)
	case http.MethodDelete:
		err:=json.NewDecoder(r.Body).Decode(&users)
		if err != nil {
			panic(err)
		}
		h.Userss.U_Delete(users.Id)
	}
}
