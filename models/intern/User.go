package intern

import (
	"encoding/base64"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
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
	u.FirstName = firstname
	u.LastName = lastname
	if password == "" {
		return u, errors.New("invalid password")
	}
	users.Insert(u)
	return u, nil
}

// AuthenticateUser checks a username/password to see if it's valid
func AuthenticateUser(username string, password string, users *mgo.Collection) string {
	user, err := GetUser(username, users)
	if err != nil {
		return ""
	}
	if password == user.Password {
		authKey := user.getAuthKey(users)
		return authKey
	}
	return ""
}

// VerifyAuthKey returns whether a username authkey pair is valid
func VerifyAuthKey(username string, key string, users *mgo.Collection) bool {
	user, err := GetUser(username, users)
	if err != nil {
		return false
	}
	compare := base64.StdEncoding.EncodeToString([]byte(user.Password))
	if compare == key {
		return true
	}
	return false
}

// User represents the MongoDB model for login/authentication
type User struct {
	ID          string              `json:"id" bson:"id"`
	Username    string              `json:"username" bson:"username"`
	FirstName   string              `json:"firstname" bson:"firstname"`
	LastName    string              `json:"lastname" bson:"lastname"`
	Password    string              `json:"password" bson:"password"`
	Description string              `json:"description" bson:"description"`
	Proposed    []string            `json:"proposed" bson:"proposed"`
	Requested   []string            `json:"requested" bson:"requested"`
	Reviews     []string            `json:"reviews" bson:"reviews"`
	AuthKeysD   [5]string           `json:"auth_key" bson:"auth_key"`
	Classes     map[string][]string `json:"classes" bson:"classes"`
}

func (u *User) getAuthKey(users *mgo.Collection) string {
	return base64.StdEncoding.EncodeToString([]byte(u.Password))
}

func GetUser(username string, users *mgo.Collection) (*User, error) {
	user := new(User)
	user.Classes = make(map[string][]string)
	err := users.Find(bson.M{"username": username}).One(&user)
	return user, err
}

func UpdateUser(user User, users *mgo.Collection) {
	users.Update(bson.M{"id": user.ID}, user)
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
