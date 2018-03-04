package intern

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Review struct {
	ID          string `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
	Value       int    `json:"value" bson:"value"`
}

func generateReviewID(reviews *mgo.Collection) string {
	count, _ := reviews.Count()
	return strconv.Itoa(count + 1)
}

func GetReview(ID string, reviews *mgo.Collection) (*Review, error) {
	review := new(Review)
	err := reviews.Find(bson.M{"id": ID}).One(&review)
	return review, err
}

func UpdateReview(review Review, reviews *mgo.Collection) {
	reviews.Update(bson.M{"id": review.ID}, review)
}
