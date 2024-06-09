package handler

import (
	"encoding/json"
	"my_mod/model"
	"net/http"
)

func (h *Handler) Problem(w http.ResponseWriter, r *http.Request) {
	var problem model.Problem
	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&problem)
		if err != nil {
			panic(err)
		}
		h.Problems.P_Create(problem)
	case http.MethodGet:
		problems, err := h.Problems.P_Read(problem)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(problems)
		if err != nil {
			panic(err)
		}
	case http.MethodPut:
		err := json.NewDecoder(r.Body).Decode(&problem)
		if err != nil {
			panic(err)
		}
		h.Problems.P_Update(problem.Id, problem.Name)
	case http.MethodDelete:
		err := json.NewDecoder(r.Body).Decode(&problem)
		if err != nil {
			panic(err)
		}
		h.Problems.P_Delete(problem.Id)
	}
}

func (h *Handler)GetBayId(w http.ResponseWriter, r *http.Request) {
var problem model.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)
		if err != nil {
			panic(err)
		}
		h.Problems.ReadById(problem.Id)
	}