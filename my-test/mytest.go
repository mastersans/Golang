package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Connect to MySQL database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		fmt.Println("Error connecting to MySQL database:", err)
		return
	}
	defer db.Close()

	// Create a Gin router
	router := gin.Default()

	// Define routes and handlers
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the Gin server
	router.Run()
}
