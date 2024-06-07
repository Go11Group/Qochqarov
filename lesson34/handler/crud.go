package handler

import (
	"encoding/json"
	"fmt"
	"my_mod/model"
	"net/http"
)

func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err:=json.NewDecoder(r.Body).Decode(&product)
	// fmt.Println(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
	h.Product.C_Create(product)
}

