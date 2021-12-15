package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hisyntax/apimal/controllers"
)

func InitRouters() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.Default()
	animal := router.Group("/animal")
	{
		animal.POST("/create", controllers.CreateAnimalHandler)
		animal.GET("/animals", controllers.GetAnimalsHandler)
		animal.GET("/:animal_id", controllers.GetAnimalHandler)
	}

	router.Run(":" + port)
}
