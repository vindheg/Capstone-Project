package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"recipe-api/internal/database"
	"recipe-api/internal/models"
)

func CreateRecipe(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	ingredients := c.PostForm("ingredients")

	file, _ := c.FormFile("image")
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	path := "uploads/" + filename

	os.MkdirAll("uploads", os.ModePerm)
	c.SaveUploadedFile(file, path)

	recipe := models.Recipe{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Ingredients: ingredients,
		ImageURL:    path,
	}

	database.DB.Create(&recipe)

	c.JSON(http.StatusCreated, recipe)
}

func GetRecipes(c *gin.Context) {
	var recipes []models.Recipe
	database.DB.Find(&recipes)
	c.JSON(http.StatusOK, recipes)
}