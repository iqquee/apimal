package main

import (
	"log"

	"github.com/iqquee/apimal/routers"
	"github.com/joho/godotenv"
)

// the init function gets called before the main function
func init() {
	//Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	//leads to the routers package
	routers.InitRouters()
}
