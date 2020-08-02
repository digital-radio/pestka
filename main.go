package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type M map[string]interface{}

func errorMarshalling(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error": "error marshalling data"}`))
}

func networkCollection(w http.ResponseWriter, r *http.Request) {
	var data = M{
		"items": 1,
	}

	var body, err = json.Marshal(data)

	if err != nil {
		errorMarshalling(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	var r = mux.NewRouter()
	r.HandleFunc("/networks", networkCollection).Methods(http.MethodGet)
	r.HandleFunc("/", notFound)

	var port = 8080
	var addr = fmt.Sprint(":", port)

	log.Println("Listening on port: ", port)
	log.Fatal(http.ListenAndServe(addr, r))
}
