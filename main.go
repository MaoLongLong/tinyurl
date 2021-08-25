package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maolonglong/tinyurl/handler"
	"github.com/maolonglong/tinyurl/storage"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)
	r.GET("/:shortUrl", handler.HandleShortUrlRedirect)

	storage.InitializeStorage()
	log.Fatal(r.Run(":9808"))
}
