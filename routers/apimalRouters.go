package routers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouters() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	routers := gin.Default()
	animal := routers.Group("/animal")
	{
		animal.POST("/create")
		animal.GET("/animal/animals")
		animal.GET("/animal/:animal_id")
		animal.GET("/animal?search=")
	}
}
