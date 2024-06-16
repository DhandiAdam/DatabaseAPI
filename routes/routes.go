package routes

import (
	"myapp/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_item.html", nil)
	})

	r.GET("/items", controllers.GetItems) // Rute untuk mendapatkan semua item
	r.POST("/items", controllers.CreateItem)

	return r
}
