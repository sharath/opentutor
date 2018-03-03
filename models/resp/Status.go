package resp

import "github.com/gin-gonic/gin"

func Error(err error) gin.H {
	return gin.H{
		"status": err,
	}
}

func OK() gin.H {
	return gin.H{
		"status": "okay, i guess",
	}
}