package handler

import (
	"encoding/json"
	"my_mod/model"
	"net/http"
)

func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&product)
		// fmt.Println(product)
		if err != nil {
			panic(err)
		}
		h.Product.C_Create(product)
	case http.MethodGet:
		pra, err := h.Product.C_Read(product)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(pra)
		if err != nil {
			panic(err)
		}
	case http.MethodPut:
		err:=json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			panic(err)
		}
		h.Product.C_Update(product.Year,product.Name)
	case http.MethodDelete:
		err:=json.NewDecoder(r.Body).Decode(product)
		if err != nil {
			panic(err)
		}
		h.Product.C_Delete(product.Year)
	}
}
