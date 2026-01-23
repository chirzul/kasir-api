package controllers

import (
	"encoding/json"
	"kasir-api/internal/models"
	"net/http"
	"strconv"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DummyCategories)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	newCategory.ID = len(models.DummyCategories) + 1
	models.DummyCategories = append(models.DummyCategories, newCategory)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.DummyCategories)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, p := range models.DummyCategories {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	for i := range models.DummyCategories {
		if models.DummyCategories[i].ID == id {
			models.DummyCategories[i] = updatedCategory
			models.DummyCategories[i].ID = id
			json.NewEncoder(w).Encode(models.DummyCategories[i])
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	for i, p := range models.DummyCategories {
		if p.ID == id {
			models.DummyCategories = append(models.DummyCategories[:i], models.DummyCategories[i+1:]...)

			json.NewEncoder(w).Encode(map[string]string{"message": "Success Delete"})
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}
