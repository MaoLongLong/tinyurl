package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maolonglong/tinyurl/shortener"
	"github.com/maolonglong/tinyurl/storage"
)

var host string

func init() {
	host = os.Getenv("HOST")
	if host == "" {
		host = "http://localhost:9808/"
	}
}

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	storage.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := storage.RetrieveInitialUrl(shortUrl)
	if initialUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "page not found",
		})
		return
	}
	c.Redirect(http.StatusFound, initialUrl)
}
