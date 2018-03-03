package tutoringApp

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
	router.GET("/api/register", register)
	router.GET("/api/login", login)
	router.Run(":8080")
}

func register(context *gin.Context) {
	u := context.PostForm("username")
	p := context.PostForm("password")
	_, err  := intern.CreateUser(u, p, database.C("users"))
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
	AuthKey, err  := intern.AuthenticateUser(u, p, database.C("users"))
	if err != nil {
		context.JSON(http.StatusBadRequest, resp.Error(err))
		return
	}
	context.JSON(200, resp.LoginResp(AuthKey))
}