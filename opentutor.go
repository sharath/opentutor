package tutoringApp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", status)
	router.GET("/api/sample", sample)
	router.Run(":8080")
}

func status(context *gin.Context) {
	context.JSON(http.StatusOK, )
}

func sample(context *gin.Context) {
	context.JSON(200, gin.H{
		"test": "test",
	})
}