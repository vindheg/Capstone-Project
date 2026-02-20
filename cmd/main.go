package main

import (
	"recipe-api/internal/database"
	"recipe-api/internal/routes"
)

func main() {
	database.Connect()
	r := routes.SetupRouter()
	r.Run(":8080")
}