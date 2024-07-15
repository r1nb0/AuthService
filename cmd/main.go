package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/r1nb0/UserService/internal/api"
	"github.com/r1nb0/UserService/internal/config"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error of loading .env file %s", err.Error())
	}
	cfg := config.GetConfig()
	serv := api.NewAppServer(cfg)
	serv.Run()
}
