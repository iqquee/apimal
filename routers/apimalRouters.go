package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iqquee/apimal/controllers"
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
		animal.PUT("/:animal_id", controllers.UpdateAnimalHandler)
		animal.DELETE("/:animal_id", controllers.DeleteAnimalHandler)
		animal.GET("/", controllers.SearchAnimalHandler)
	}

	router.Run(":" + port)
}
