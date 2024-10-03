package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var products = map[string]Product{
	"1": {ID: "1", Name: "Product A"},
	"2": {ID: "2", Name: "Product B"},
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product, exists := products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func main() {
	http.HandleFunc("/product", getProduct)
	http.ListenAndServe(":8085", nil)
}
