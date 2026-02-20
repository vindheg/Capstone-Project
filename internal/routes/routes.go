package routes

import (
	"github.com/gin-gonic/gin"
	"recipe-api/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/recipes", handlers.CreateRecipe)
	r.GET("/recipes", handlers.GetRecipes)

	return r
}