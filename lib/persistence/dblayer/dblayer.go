package dblayer

import (
	"github.com/golabay/refs-ebook-service/lib/persistence"
	"github.com/golabay/refs-ebook-service/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
)

func NewPersistenceLayer(option DBTYPE, connect string) (persistence.DatabaseHandler, error) {
	switch option {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connect)
	}
	return nil, nil
}
