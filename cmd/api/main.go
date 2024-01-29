package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tbrito/ms-lanchonetedarua-clientes/cmd/api/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	controllers.ClienteRoutes(r)

	r.Run()
}
