package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/sharath/opentutor/models/intern"
)

func FindTutorResp(subject string, class string, users *mgo.Collection) gin.H {
	// this should use the aggregated pipline in the future
	var all []intern.User
	users.Find(nil).All(&users)
	var found []intern.User
	for _, u := range all {
		for _, s := range u.Classes[subject] {
			if s == class {
				found = append(found, u)
			}
		}
	}
	return gin.H{
		"tutors": found,
	}
}
