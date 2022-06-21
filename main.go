package main

import (
	"log"
	"net/http"
	"webhooker/src/handlers"
	_ "webhooker/src/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "OK",
		})
	})

	r.POST("/", func(c *gin.Context) {
		err := handlers.Do(c)
		if err != nil {
			log.Printf("handles error %+v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": gin.H{
					"status_code": http.StatusBadRequest,
					"status":      "bad request",
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": gin.H{
					"status_code": http.StatusOK,
					"status":      "ok",
				},
				"data": "ok",
			})
		}

	})

	r.POST("/log/minio-social/", func(c *gin.Context) {
		err := handlers.LogRecodes(c)
		if err != nil {
			log.Printf("handles error %+v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": gin.H{
					"status_code": http.StatusBadRequest,
					"status":      "bad request",
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": gin.H{
					"status_code": http.StatusOK,
					"status":      "ok",
				},
				"data": "ok",
			})
		}

	})

	r.Run()
}
