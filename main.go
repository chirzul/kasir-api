package main

import (
	"fmt"
	"kasir-api/internal/router"
	"net/http"
)

func main() {
	mux := router.SetupRouter()

	fmt.Println("Server running at :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error running server")
	}
}
