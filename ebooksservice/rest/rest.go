package rest

import (
	"github.com/golabay/refs-ebook-service/lib/persistence"
	"github.com/gorilla/mux"
	"net/http"
)

func ServeAPI(endpoint string, dbhandler persistence.DatabaseHandler) error {
	handler := NewEbookHandler(dbhandler)
	r := mux.NewRouter()
	ebooksrouter := r.PathPrefix("/ebooks").Subrouter()
	ebooksrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEbookHandler)
	ebooksrouter.Methods("POST").Path("").HandlerFunc(handler.NewEbookHandler)
	return http.ListenAndServe(endpoint, r)
}
