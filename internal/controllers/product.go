package controllers

import (
	"encoding/json"
	"kasir-api/internal/models"
	"net/http"
	"strconv"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DummyProducts)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	newProduct.ID = len(models.DummyProducts) + 1
	models.DummyProducts = append(models.DummyProducts, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.DummyProducts)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range models.DummyProducts {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	for i := range models.DummyProducts {
		if models.DummyProducts[i].ID == id {
			models.DummyProducts[i] = updatedProduct
			models.DummyProducts[i].ID = id
			json.NewEncoder(w).Encode(models.DummyProducts[i])
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	for i, p := range models.DummyProducts {
		if p.ID == id {
			models.DummyProducts = append(models.DummyProducts[:i], models.DummyProducts[i+1:]...)

			json.NewEncoder(w).Encode(map[string]string{"message": "Success Delete"})
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}
