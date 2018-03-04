package resp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sharath/opentutor/models/intern"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Requested(usr string, users *mgo.Collection) gin.H {
	var user intern.User
	users.Find(bson.M{"username": usr}).One(&user)
	type MinInfo struct {
		FirstName string `json:"firstname" bson:"firstname"`
		LastName  string `json:"lastname" bson:"lastname"`
	}
	t := make(map[string]MinInfo)
	for _, id := range user.Requested {
		var sqry intern.User
		users.Find(bson.M{"id": id}).All(&sqry)
		var min MinInfo
		clean, _ := json.Marshal(sqry)
		json.Unmarshal(clean, &min)
		t[sqry.Username] = min
	}
	return gin.H{
		"requested": t,
	}
}
