package main

import (
	"context"
	_ "kasir-api/docs"
	"kasir-api/internal/config"
	"kasir-api/internal/database"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/router"
	"kasir-api/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title			Kasir API
// @version		1.0
// @description	This is API documentation for kasir-api
// @BasePath		/api/v1
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stdout,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()

	pool, err := database.InitPostgresDB(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	defer pool.Close()

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	repository := repository.New(pool)

	productService := service.NewProductService(repository)
	productHandler := handler.NewProductHandler(productService)

	categoryService := service.NewCategoryService(repository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	handlers := &router.Handlers{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
	}
	router.RegisterRoutes(app, handlers)

	go func() {
		log.Info().
			Str("port", cfg.AppPort).
			Msg("HTTP server started")

		if err := app.Listen(":" + cfg.AppPort); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("server crashed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutdown signal received")

	if err := app.Shutdown(); err != nil {
		log.Error().Err(err).Msg("graceful shutdown failed")
	} else {
		log.Info().Msg("server stopped gracefully")
	}
}
