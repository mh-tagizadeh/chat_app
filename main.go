package main


import (
    "github.com/gin-gonic/gin"

	"chat_app/database"
	"chat_app/routes"
	"os"
)


var router = gin.Default()

func main(){


	// Connect with database
	// @parameter connection_string => Database details for connect postgres database
	database.Connect(os.Getenv("DATABASE_CONNECTION"))

	// Migrate tables 
	database.Migrates(database.Options{
		Migrate: true,
		DB: database.Connection,
	})

	getRoutes()

	router.Run("localhost:8080")
}


func getRoutes() {
	v1 := router.Group("/v1")

	routes.AddAuthGroup(v1)
}

