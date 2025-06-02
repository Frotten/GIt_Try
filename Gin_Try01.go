package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Owl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ //gin.H是一个map类型，键的类型是string，值的类型为空接口，可以用于简化 JSON 数据的创建和处理
		"message": "My Stupid and cute Owl",
	})
}
func main() {
	r := gin.Default() //这里是返回默认的路由引擎
	r.GET("/Owl", Owl) //指定用户使用GET访问/Owl路径时，调用Owl函数
	r.PUT("/Owl", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Put": "This is a PUT request",
		})
	})
	r.POST("/Owl", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Post": "This is a POST request",
		})
	})
	r.DELETE("/Owl", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Delete": "This is a DELETE request",
		})
	})
	r.Run(":9090") //启动服务，默认端口是8080
}
