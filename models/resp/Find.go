package resp

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func FindTutorResp(subject string, class string, users *mgo.Collection) gin.H {
	return gin.H{
		"message": "to be implemented",
	}
}
