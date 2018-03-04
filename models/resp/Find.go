package resp

import (
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)

func FindTutorResp(subject string, class string, users *mgo.Collection) gin.H {
	return gin.H{
		"message": "to be implemented",
	}
}