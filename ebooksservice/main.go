package main

import (
	"github.com/golabay/refs-ebook-service/ebooksservice/rest"
	"github.com/golabay/refs-ebook-service/lib/configuration"
	"github.com/golabay/refs-ebook-service/lib/persistence/dblayer"
	"log"
)

func main() {
	conf := configuration.ExtractConfiguration()

	dbhandler, _ := dblayer.NewPersistenceLayer(conf.DatabaseType, conf.DBConnection)

	log.Fatal(rest.ServeAPI(conf.RestfulEndpoint, dbhandler))
}
