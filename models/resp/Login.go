package resp

import "github.com/gin-gonic/gin"

func LoginResp(AuthKey string) gin.H {
	return gin.H{
		"auth_key": AuthKey,
	}
}