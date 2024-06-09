package handler

import (
	"encoding/json"
	"my_mod/model"
	"net/http"
)

func (h *Handler) Solved_problems(w http.ResponseWriter, r *http.Request) {
	var solved model.Solved_problems
	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&solved)
		if err != nil {
			panic(err)
		}
		h.Solvedproblems.S_Create(solved)
	case http.MethodGet:
		solveds, err := h.Solvedproblems.S_Read(solved)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(solveds)
		if err != nil {
			panic(err)
		}
	case http.MethodPut:
		err := json.NewDecoder(r.Body).Decode(&solved)
		if err != nil {
			panic(err)
		}
		h.Solvedproblems.S_Update(solved.Id, solved.Name)
	case http.MethodDelete:
		err := json.NewDecoder(r.Body).Decode(&solved)
		if err != nil {
			panic(err)
		}
		h.Solvedproblems.S_Delete(solved.Id)
	}
}
