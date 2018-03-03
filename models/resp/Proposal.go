package resp

import (
	"github.com/sharath/opentutor/models/intern"
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func Proposal(usr string, users *mgo.Collection) gin.H {
	var user intern.User
	users.Find(bson.M{"username":usr}).One(&user)
	type MinInfo struct {
		FirstName string `json:"firstname" bson:"firstname"`
		LastName  string `json:"lastname" bson:"lastname"`
	}
	t := make(map[string]MinInfo)
	for _, id := range user.Proposed {
		var sqry intern.User
		users.Find(bson.M{"id": id}).All(&sqry)
		var min MinInfo
		clean, _ := json.Marshal(sqry)
		json.Unmarshal(clean, &min)
		t[sqry.Username] = min
	}
	return gin.H{
		"proposed": t,
	}
}
