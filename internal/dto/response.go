package dto

type ErrorResponse struct {
	StatusCode int     `json:"statusCode"`
	StatusDesc string  `json:"statusDesc"`
	Error      *string `json:"error"`
}

type SuccessResponse struct {
	StatusCode int    `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
	Data       any    `json:"data,omitempty"`
}

type ProductResponse struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"categoryName"`
	CategoryDesc string  `json:"categoryDesc"`
}

type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
