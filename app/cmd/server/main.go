package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	router.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	})

	router.GET("/api/v1/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "project-atlas",
			"environment": "development",
			"status": "running",
		})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Freshair API is running",
    })
})

	router.Run(":8080")
}