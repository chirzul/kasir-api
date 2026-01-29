package main

import (
	"context"
	"fmt"
	"kasir-api/internal/config"
	"kasir-api/internal/database"
	"kasir-api/internal/router"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.InitPostgresDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	mux := router.SetupRouter(pool)

	fmt.Println("Server running at :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error running server")
	}
}
