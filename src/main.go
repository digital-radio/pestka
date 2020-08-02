package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MonkeyBuisness/golang-iwlist"
	"github.com/gorilla/mux"
)

var interfaceName = "wlan0"

func response(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}

	w.WriteHeader(status)
	w.Write(body)
}

func handleError(w http.ResponseWriter, err error) {
	log.Println(err)
	var data = map[string]string{
		"message": err.Error(),
	}
	response(w, data, http.StatusInternalServerError)
}

func networkCollection(w http.ResponseWriter, r *http.Request) {
	cells, err := wlist.Scan(interfaceName)
	if err != nil {
		handleError(w, err)
		return
	}

	response(w, cells, http.StatusOK)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"message": "not found",
	}
	response(w, data, http.StatusOK)
}

func main() {
	var r = mux.NewRouter()
	r.HandleFunc("/networks", networkCollection).Methods(http.MethodGet)
	r.HandleFunc("/", notFound)

	var port = 8080
	var addr = fmt.Sprint(":", port)

	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(addr, r))
}
