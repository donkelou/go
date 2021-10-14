package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/form_post",func(c *gin.Context){
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick","anonymous")
		c.JSON(200,gin.H{
			"status":"posted",
			"message":message,
			"nick":nick,
		})
	})
	router.Run(":8080")

}
