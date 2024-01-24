package usecase

import (
	"log"
	"net/http"

	"example.com/server/config"
	"example.com/server/models"
	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item

	rows, err := config.Db.Query("SELECT * FROM items")
	if err != nil {
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb models.Item
		if err := rows.Scan(&alb.ID, &alb.Name, &alb.Price); err != nil {
			return
		}
		items = append(items, alb)
	}
	if err := rows.Err(); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}
func GetItemById(c *gin.Context) {
	var items []models.Item
	id := c.Param("id")
	rows, err := config.Db.Query("SELECT * FROM items WHERE id = ?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		var alb models.Item
		if err := rows.Scan(&alb.ID, &alb.Name, &alb.Price); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		items = append(items, alb)
	}
	if err := rows.Err(); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}
func PostItems(c *gin.Context) {
	// var items []Item
	var newItem models.Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	_, err := config.Db.Query("INSERT INTO items (name, price) VALUES (?,?)", newItem.Name, newItem.Price)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, "successfully added")

}
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem models.Item
	if err := c.BindJSON(&updatedItem); err != nil {
		return
	}

	_, err := config.Db.Query("UPDATE items SET name = ?, price = ? WHERE id = ?", updatedItem.Name, updatedItem.Price, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, updatedItem)
}
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	_, err := config.Db.Query("DELETE from items WHERE id = ?", id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, "deleted")
}
