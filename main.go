package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Plus(c *gin.Context) {
	a, err := strconv.Atoi(c.Query("a"))
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid parameter",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": a + b,
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/plus", Plus)

	r.Run()
}
