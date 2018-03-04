package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sharath/opentutor/models/intern"
	"gopkg.in/mgo.v2"
)

func Reviews(usr string, users *mgo.Collection, reviews *mgo.Collection) gin.H {
	user, _ := intern.GetUser(usr, users)
	var qry []*intern.Review
	for _, id := range user.Reviews {
		find, _ := intern.GetReview(id, reviews)
		qry = append(qry, find)
	}
	return gin.H{
		"reviews": qry,
	}
}
