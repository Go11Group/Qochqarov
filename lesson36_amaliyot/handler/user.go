package handler

import (
	"encoding/json"
	"my_mod/model"
	"net/http"
	"strconv"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	var users model.Users
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		panic(err)
	}

	h.Userss.U_Read(users, id)
}
