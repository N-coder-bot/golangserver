package main

import (
	"log"

	"example.com/server/config"
	"example.com/server/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// _ "github.com/go-sql-driver/mysql"
)

// // slice of []item --> this was done when i was using local memory.
// var items = []Item{
// 	{ID: "1", Name: "Fruit", Price: 25.3},
// 	{ID: "2", Name: "Clothe", Price: 45.3},
// 	{ID: "3", Name: "Medicine", Price: 25.3},
// }

func main() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the dotenv file: %s", err)
	}
	config.Connection()
	router := gin.Default()
	router.GET("/items", usecase.GetItems)
	router.GET("/items/:id", usecase.GetItemById)
	router.POST("/items", usecase.PostItems)
	router.PUT("/items/:id", usecase.UpdateItem)
	router.DELETE("/items/:id", usecase.DeleteItem)
	router.Run("localhost:8080")
}
