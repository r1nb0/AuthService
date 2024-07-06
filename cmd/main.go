package main

import (
	"AuthService/internal/config"
	"AuthService/internal/controller"
	"AuthService/internal/infra"
	"AuthService/internal/usecase"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	cfg, err := config.GetConfig("./configs", "config-prod", "yaml")
	if err != nil {
		log.Fatalf("error of loading configs: %v", err)
	}

	db, err := infra.InitPostgres(cfg)
	if err != nil {
		log.Fatalf("error of initializing db: %v", err)
	}

	repo := infra.NewUserRepository(db)
	uc := usecase.NewAuthService(repo, cfg)
	handler := controller.NewAuthController(uc)

	app := fiber.New()
	api := app.Group("/api")
	auth := api.Group("/auth")
	{
		auth.Post("/sign-in", handler.SignIn)
		auth.Post("/sign-up", handler.SignUp)
	}

	if err = app.Listen(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		log.Fatalf("error of starting app: %v", err)
	}
}
