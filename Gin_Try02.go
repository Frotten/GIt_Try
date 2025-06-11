package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/web", func(c *gin.Context) {
		//NameGinTry02 := c.Query("name") //这里的key为name，因此我们在对应网址后加上对应的key和value即可传入，如xxxx/web?name=AAAA
		//NameGinTry02 := c.DefaultQuery("name", "no thing to search")
		////DefaultQuery会在我们使用对应键值并且查询成功之后使用对应的键值，如果我们使用的键值与设定的键值不一样，那么就会自动输出我们的defaultValue
		NameGinTry02, ok := c.GetQuery("name") //使用GetQuery的整体效果类似于DefaultQuery，但是是要根据返回的bool类型参数来判断
		if !ok {
			NameGinTry02 = "no thing to search"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": NameGinTry02, //这里是json的输出格式，如果按照上面的web?name=AAAA的话，这里也就是按照"name": AAAA的模板来进行的
		})
	})
	_ = engine.Run(":9090")
}
