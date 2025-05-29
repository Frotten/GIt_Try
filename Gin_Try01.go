package main

import "github.com/gin-gonic/gin"

func Owl(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "My Stupid and cute Owl",
	})
}
func main() {
	r := gin.Default() //这里是返回默认的路由引擎
	r.GET("/Owl", Owl) //指定用户使用GET访问/Owl路径时，调用Owl函数
	r.Run(":9090")     //启动服务，默认端口是8080
}
