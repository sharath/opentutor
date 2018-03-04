package intern

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"encoding/base64"
	"fmt"
)

func CreateUser(username string, password string, firstname string, lastname string, users *mgo.Collection) (*User, error) {
	u := new(User)
	if !validNewUsername(users, username) {
		return u, errors.New("invalid username")
	}
	if !validPassword(password) {
		return u, errors.New("invalid password")
	}
	u.ID = generateUserID(users)
	u.Username = username
	u.Password = password
	fmt.Println(password)
	u.FirstName = firstname
	u.LastName = lastname
	if password == "" {
		return u, errors.New("invalid password")
	}
	users.Insert(u)
	return u, nil
}

// AuthenticateUser checks a username/password to see if it's valid
func AuthenticateUser(username string, password string, users *mgo.Collection) (string) {
	var user User
	users.Find(bson.M{"username": username}).One(&user)
	if password == user.Password {
		authKey := user.getAuthKey(users)
		return authKey
	}
	fmt.Println(user.Password, password)
	return ""
}

// VerifyAuthKey returns whether a username authkey pair is valid
func VerifyAuthKey(user string, key string, users *mgo.Collection) bool {
	var u User
	err := users.Find(bson.M{"username": user}).One(&u)
	if err != nil {
		return false
	}
	compare := base64.StdEncoding.EncodeToString([]byte(u.Password))
	if compare == key {
		return true
	}
	fmt.Println(compare, key)
	return false
}

// User represents the MongoDB model for login/authentication
type User struct {
	ID        string    `json:"id" bson:"id"`
	Username  string    `json:"username" bson:"username"`
	FirstName string    `json:"firstname" bson:"firstname"`
	LastName  string    `json:"lastname" bson:"lastname"`
	Password  string    `json:"password" bson:"password"`
	Proposed  []string  `json:"proposed" bson:"proposed"`
	Requested []string  `json:"requested" bson:"requested"`
	AuthKeysD [5]string `json:"auth_key" bson:"auth_key"`
}

func (u *User) getAuthKey(users *mgo.Collection) (string) {
	return base64.StdEncoding.EncodeToString([]byte(u.Password))
}

func generateUserID(users *mgo.Collection) string {
	count, _ := users.Count()
	return strconv.Itoa(count + 1)
}

func validNewUsername(users *mgo.Collection, username string) bool {
	count, _ := users.Find(bson.M{"username": username}).Count()
	if count != 0 {
		return false
	}
	return true
}

func validPassword(password string) bool {
	return !(len(password) < 7)
}
