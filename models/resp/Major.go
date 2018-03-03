package resp

import "gopkg.in/mgo.v2"

func Major(classes *mgo.Collection) map[string][]string {
	majors := make(map[string][]string)
	classes.Find(nil).All(&majors)
	return majors
}