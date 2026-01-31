package dto

type AddProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryId string  `json:"categoryId"`
}

type UpdateProductRequest struct {
	ID         string
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryId string  `json:"categoryId"`
}

type AddCategoryRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	ID          string
	Name        string `json:"name"`
	Description string `json:"description"`
}
