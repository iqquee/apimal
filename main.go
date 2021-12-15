package main

import (
	"log"

	"github.com/hisyntax/apimal/routers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	routers.InitRouters()
}
