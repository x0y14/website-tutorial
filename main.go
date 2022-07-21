package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	version = "1"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": version,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
