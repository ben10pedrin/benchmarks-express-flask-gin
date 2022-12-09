package main

import "github.com/gin-gonic/gin"

type Request struct {
	Hello string `json:"hello"`
}

func main() {
	// Running in production mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Text - GET
	router.GET("/text", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	// JSON - GET
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	// JSON - POST
	router.POST("/json", func(c *gin.Context) {
		var request Request
		c.ShouldBindJSON(&request)

		c.JSON(200, gin.H{
			"hello": request.Hello,
		})
	})

	router.Run(":3000")
}

