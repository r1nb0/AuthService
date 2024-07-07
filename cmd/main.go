package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/r1nb0/UserService/api"
	"github.com/r1nb0/UserService/configs"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error of loading .env file %s", err.Error())
	}
	serv := api.NewAppServer(configs.GetConfig())
	serv.Run()
}
