package controllers

import (
	"myapp/database"
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO items (name, description, harga) VALUES (?, ?, ?)`
	result, err := database.DB.Exec(query, newItem.Name, newItem.Description, newItem.Harga)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newItem.ID = int(lastInsertID)
	c.JSON(http.StatusOK, newItem)
}

func GetItems(c *gin.Context) {
	var items []models.Item

	rows, err := database.DB.Query("SELECT id, name, description, harga FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Harga); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}
