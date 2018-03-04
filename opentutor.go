package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
	"github.com/sharath/opentutor/models/intern"
	"github.com/sharath/opentutor/models/resp"
	"errors"
)

var database *mgo.Database

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	database = session.DB("ot")

	router := gin.Default()
	router.GET("/", status)
	router.GET("/api/user/proposed", tutorProposed)
	router.GET("/api/user/requested", tutorRequested)
	router.GET("/api/tutor", findTutor)
	router.GET("/api/classes", major)
	router.POST("/api/register", register)
	router.POST("/api/login", login)
	router.Run(":8080")
}

func register(context *gin.Context) {
	u := context.PostForm("username")
	p := context.PostForm("password")
	fn := context.PostForm("firstname")
	ln := context.PostForm("lastname")
	_, err := intern.CreateUser(u, p, fn, ln, database.C("users"))
	if err != nil {
		context.JSON(http.StatusBadRequest, resp.Error(err))
		return
	}
	context.JSON(http.StatusOK, resp.OK())
}

func status(context *gin.Context) {
	context.JSON(http.StatusOK, resp.OK())
}

func login(context *gin.Context) {
	u := context.PostForm("username")
	p := context.PostForm("password")
	authkey := intern.AuthenticateUser(u, p, database.C("users"))
	if authkey != "" {
		context.JSON(200, resp.Login(authkey))
		return
	}
	context.JSON(http.StatusBadRequest, resp.Error(errors.New("invalid login")))
}

func tutorProposed(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid  := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		context.JSON(http.StatusBadRequest, resp.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, resp.Proposal(usr, database.C("users")))
}

func tutorRequested(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid  := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		fmt.Println(usr,key)
		context.JSON(http.StatusBadRequest, resp.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, resp.Requested(usr, database.C("users")))
}

func major(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	valid  := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		fmt.Println(usr,key)
		context.JSON(http.StatusBadRequest, resp.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, resp.Major(database.C("classes")))
}

func findTutor(context *gin.Context) {
	usr := context.Query("username")
	key := context.Query("auth_key")
	subject := context.Query("subject")
	number := context.Query("class")
	valid := intern.VerifyAuthKey(usr, key, database.C("users"))
	if !valid {
		context.JSON(http.StatusBadRequest, resp.Error(errors.New("invalid login")))
		return
	}
	context.JSON(http.StatusOK, resp.FindTutorResp(subject, number, database.C("users")))
}