package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.Get("/", func(c *gin.Context){
		c.String(200,"值%v","你好gin")
	})

	r.Get("/news", func(c *gin.Context){
		c.String(200,"我是新闻页面")
	})

	r.Run(":8000")
}