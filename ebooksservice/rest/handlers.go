package rest

import (
	"encoding/json"
	"fmt"
	"github.com/golabay/refs-ebook-service/lib/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type ebookServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func NewEbookHandler(databasehandler persistence.DatabaseHandler) *ebookServiceHandler {
	return &ebookServiceHandler{
		dbhandler: databasehandler,
	}
}

func (eh *ebookServiceHandler) NewEbookHandler(w http.ResponseWriter, r *http.Request) {
	ebook := persistence.Ebook{}
	err := json.NewDecoder(r.Body).Decode(&ebook)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while decoding event data %s"}`, err)
		return
	}
	id, err := eh.dbhandler.AddEbook(ebook)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while persisting event %d %s"}`, id, err)
		return
	}
	fmt.Fprintf(w, `{"id",%d}`, id)
}

func (eh *ebookServiceHandler) FindEbookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search criteria found, you can either
						search by id via /id/4
						to search by name via /name/coldplayconcert"}`)
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either search
						by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}
	var ebook persistence.Ebook
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		ebook, err = eh.dbhandler.FindEbookByName(searchkey)
	}
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&ebook)
}
