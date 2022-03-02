package main

import "github.com/gin-gonic/gin"

func main() {
	c := gin.Default()
	c.GET("/", func(c *gin.Context) {
		c.JSON(200, "我运行在linux上!!")

	})

	c.Run()
}
