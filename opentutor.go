package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sharath/opentutor/controllers"
	"github.com/sharath/opentutor/models/intern"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
)

var database *mgo.Database

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	database = session.DB("ot")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", status)
	router.GET("/api/user/proposed", userProposed)
	router.GET("/api/user/requested", userRequested)
	router.GET("/api/user/reviews", userReviews)
	router.GET("/api/tutor", findTutor)
	router.GET("/api/classes", major)
	router.POST("/api/register", register)
	router.POST("/api/login", login)
	router.Run(":80")
}

func register(context *gin.Context) {
	u := context.PostForm("username")
	p := context.PostForm("password")
	fn := context.PostForm("firstname")
	ln := context.PostForm("lastname")
	_, err := intern.CreateUser(u, p, fn, ln, database.C("users"))
	if err != nil {
		context.JSON(http.StatusBadRequest, controllers.Error(err))
		return
	}
	context.JSON(http.StatusOK, controllers.OK())
}

func status(context *gin.Context) {
	context.JSON(http.StatusOK, controllers.OK())
}

func login(context *gin.Context) {
	u := context.PostForm("username")
	p := context.PostForm("password")
	authkey := intern.AuthenticateUser(u, p, database.C("users"))
	if authkey != "" {
		context.JSON(200, controllers.Login(authkey))
		return
	}
	context.JSON(http.StatusBadRequest, controllers.Error(errors.New("invalid login")))
}

func userProposed(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		context.JSON(http.StatusBadRequest, controllers.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, controllers.Proposal(usr, database.C("users")))
}

func userRequested(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		fmt.Println(usr, key)
		context.JSON(http.StatusBadRequest, controllers.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, controllers.Requested(usr, database.C("users")))
}

func userReviews(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		fmt.Println(usr, key)
		context.JSON(http.StatusBadRequest, controllers.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, controllers.Reviews(usr, database.C("users"), database.C("reviews")))
}

func major(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		fmt.Println(usr, key)
		context.JSON(http.StatusBadRequest, controllers.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, controllers.Major(database.C("classes")))
}

func findTutor(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	subject := context.Query("subject")
	number := context.Query("class")
	valid := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		context.JSON(http.StatusBadRequest, controllers.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, controllers.FindTutorResp(subject, number, database.C("users")))
}
