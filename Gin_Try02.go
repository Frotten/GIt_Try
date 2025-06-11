package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/web", func(c *gin.Context) {
		getname := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"name": getname,
		})
	})
	engine.Run(":9090")
}
