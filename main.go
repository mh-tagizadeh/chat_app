package main


import (
    "github.com/gin-gonic/gin"

	"github.com/mh-tagizadeh/chat_app/database"
	"os"
)

func main(){


	// Connect with database
	// @parameter connection_string => Database details for connect postgres database
	database.Connect(os.Getenv("DATABASE_CONNECTION"))

	// Migrate tables 
	database.Migrates(Options{
		Migrate: true,
		DB: database.Connection,
	})

    router := gin.Default()

	router.Run("localhost:8080")
}
