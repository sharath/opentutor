package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", sample)
	router.Run(":8080")
}

func sample(context *gin.Context) {
	context.JSON(200, gin.H{
		"test": "test",
	})
}