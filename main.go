package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.Get("/", func(c *gin.Context){
		c.String(200,"值%v","你好gin")
	})

	r.Run()
}