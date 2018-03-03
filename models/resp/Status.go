package resp

import "github.com/gin-gonic/gin"

func Error(err error) gin.H {
	return gin.H{
		"error": err,
	}
}