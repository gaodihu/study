package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
	router.Run(":8000")
}
