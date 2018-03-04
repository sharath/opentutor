package resp

import (
	"gopkg.in/mgo.v2"
)

func Major(classes *mgo.Collection) map[string][]string {
	t := make(map[string][]string)
	classes.Find(nil).One(&t)
	return t
}