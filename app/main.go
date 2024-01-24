package main

import (
	"database/sql"
	"log"

	"example.com/server/config"
	"example.com/server/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// _ "github.com/go-sql-driver/mysql"
)

// // slice of []item --> this was done when i was using local memory.
//
//	var items = []Item{
//		{ID: "1", Name: "Fruit", Price: 25.3},
//		{ID: "2", Name: "Clothe", Price: 45.3},
//		{ID: "3", Name: "Medicine", Price: 25.3},
//	}
type ItemService struct {
	db *sql.DB
}

func NewItemService(db *sql.DB) *ItemService {
	return &ItemService{db: db}
}
func (it *ItemService) getItemById(c *gin.Context) {
	usecase.GetItemById(c)
}
func (it *ItemService) getItems(c *gin.Context) {
	usecase.GetItems(c)
}
func (it *ItemService) postItems(c *gin.Context) {
	usecase.PostItems(c)
}
func (it *ItemService) updateItem(c *gin.Context) {
	usecase.UpdateItem(c)
}
func (it *ItemService) deleteItem(c *gin.Context) {
	usecase.DeleteItem(c)
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the dotenv file: %s", err)
	}

	var db *sql.DB = config.Connection()

	itemService := NewItemService(db)

	router := gin.Default()
	router.GET("/items", itemService.getItems)
	router.GET("/items/:id", itemService.getItemById)
	router.POST("/items", itemService.postItems)
	router.PUT("/items/:id", itemService.updateItem)
	router.DELETE("/items/:id", itemService.deleteItem)
	router.Run("localhost:8080")
}
