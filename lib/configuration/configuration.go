package configuration

import "github.com/golabay/refs-ebook-service/lib/persistence/dblayer"

var (
	DBType = dblayer.DBTYPE("mongodb")
	DBConn = "127.0.0.1"
	RestEP = "localhost:8080"
)

type ServiceConfig struct {
	DatabaseType    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulendpoint"`
}

func ExtractConfiguration() ServiceConfig {
	conf := ServiceConfig{
		DBType,
		DBConn,
		RestEP,
	}
	return conf
}
