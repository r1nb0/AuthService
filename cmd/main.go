package main

import (
	"AuthService/internal/config"
	"AuthService/internal/controllers"
	"AuthService/internal/infra"
	"AuthService/internal/middleware"
	"AuthService/internal/usecase"
	"AuthService/pkg/utils"
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

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error of loading configs: %s", err.Error())
	}

	db, err := infra.InitPostgres(cfg)
	if err != nil {
		log.Fatalf("error of initializing db: %s", err.Error())
	}

	userRepo := infra.NewUserRepository(db)
	jwtUtil := utils.NewJWTUtil(cfg)
	authMiddleware := middleware.NewAuthMiddleware(jwtUtil)
	authUsecase := usecase.NewAuthService(userRepo, jwtUtil, cfg)
	userUsecase := usecase.NewUserService(userRepo)
	authController := controllers.NewAuthController(authUsecase)
	userController := controllers.NewUserController(userUsecase)

	app := fiber.New()
	api := app.Group("/api")
	auth := api.Group("/auth")

	{
		auth.Post("/sign-in", authController.SignIn)
		auth.Post("/sign-up", authController.SignUp)
	}

	users := api.Group("/users")
	{
		users.Get("/", userController.GetAll)
		users.Get("/:id", userController.GetByID)
		users.Put("/", userController.Update, authMiddleware.Authentication)
	}

	if err = app.Listen(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		log.Fatalf("error of starting app: %s", err.Error())
	}
}
