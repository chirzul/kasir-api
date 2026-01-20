package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var product = []Product{
	{
		ID:    1,
		Name:  "Indomie",
		Price: 3500,
		Stock: 10,
	},
	{
		ID:    2,
		Name:  "Aqua",
		Price: 3000,
		Stock: 40,
	},
	{
		ID:    3,
		Name:  "Kecap",
		Price: 3000,
		Stock: 40,
	},
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
		} else if r.Method == "POST" {
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
			}

			newProduct.ID = len(product) + 1
			product = append(product, newProduct)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
		}
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetProductByID(w, r)
		}
		if r.Method == "PUT" {
			UpdateProductByID(w, r)
		}
		if r.Method == "DELETE" {
			DeleteProductByID(w, r)
		}
	})

	fmt.Println("Server running at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error running server")
	}
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range product {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	for i := range product {
		if product[i].ID == id {
			w.Header().Set("Content-Type", "application/json")
			product[i] = updatedProduct
			product[i].ID = id
			json.NewEncoder(w).Encode(product[i])
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	for i, p := range product {
		if p.ID == id {
			product = append(product[:i], product[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Success Delete"})
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}
