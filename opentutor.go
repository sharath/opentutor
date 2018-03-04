package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
	"github.com/sharath/opentutor/models/intern"
	"github.com/sharath/opentutor/models/resp"
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
	AuthKey, err := intern.AuthenticateUser(u, p, database.C("users"))
	if err != nil {
		context.JSON(http.StatusBadRequest, resp.Error(err))
		return
	}
	context.JSON(200, resp.Login(AuthKey))
}

func tutorProposed(context *gin.Context) {
	usr := context.GetString("username")
	key := context.GetString("auth_key")
	valid, err := intern.VerifyAuthKey(usr, key, database.C("users"))
	if err != nil || !valid {
		context.JSON(http.StatusBadRequest, resp.Error(err))
		return
	}
	context.JSON(http.StatusOK, resp.Proposal(usr, database.C("users")))
}

func tutorRequested(context *gin.Context) {
	usr := context.GetString("username")
	key := context.GetString("auth_key")
	valid, err := intern.VerifyAuthKey(usr, key, database.C("users"))
	if err != nil || !valid {
		context.JSON(http.StatusBadRequest, resp.Error(err))
		return
	}
	context.JSON(http.StatusOK, resp.Requested(usr, database.C("users")))
}

func major(context *gin.Context) {
	usr := context.GetString("username")
	key := context.GetString("password")
	valid, err := intern.VerifyAuthKey(usr, key, database.C("users"))
	if err != nil || !valid {
		context.JSON(http.StatusBadRequest, resp.Error(err))
		return
	}
	context.JSON(http.StatusOK, resp.Major(database.C("classes")))
}