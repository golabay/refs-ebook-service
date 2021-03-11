package persistence

import "gopkg.in/mgo.v2/bson"

type Ebook struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
	Pages int           `json:"pages" bson:"pages"`
}
