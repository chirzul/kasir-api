package controllers

import (
	"encoding/json"
	"kasir-api/internal/models"
	"net/http"
	"strconv"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DummyProduct)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	newProduct.ID = len(models.DummyProduct) + 1
	models.DummyProduct = append(models.DummyProduct, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.DummyProduct)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range models.DummyProduct {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	for i := range models.DummyProduct {
		if models.DummyProduct[i].ID == id {
			w.Header().Set("Content-Type", "application/json")
			models.DummyProduct[i] = updatedProduct
			models.DummyProduct[i].ID = id
			json.NewEncoder(w).Encode(models.DummyProduct[i])
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	for i, p := range models.DummyProduct {
		if p.ID == id {
			models.DummyProduct = append(models.DummyProduct[:i], models.DummyProduct[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Success Delete"})
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}
