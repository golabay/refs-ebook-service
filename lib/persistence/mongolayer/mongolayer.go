package mongolayer

import (
	"github.com/golabay/refs-ebook-service/lib/persistence"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB     = "mekong"
	EBOOKS = "ebooks"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

func NewMongoDBLayer(connect string) (persistence.DatabaseHandler, error) {
	s, err := mgo.Dial(connect)
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) AddEbook(e persistence.Ebook) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	return []byte(e.ID), s.DB(DB).C(EBOOKS).Insert(e)
}

func (mgoLayer *MongoDBLayer) FindEbookByName(title string) (persistence.Ebook, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Ebook{}
	err := s.DB(DB).C(EBOOKS).Find(bson.M{"title": title}).One(&e)
	return e, err
}
