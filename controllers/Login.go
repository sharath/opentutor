package controllers

import "github.com/gin-gonic/gin"

func Login(AuthKey string) gin.H {
	return gin.H{
		"auth_key": AuthKey,
	}
}
