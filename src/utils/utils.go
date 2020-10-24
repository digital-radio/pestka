//Package utils consists of small helper functions.
package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/digital-radio/pestka/src/app/exceptions"
)

//HandleError sends 500 response with error message in the body.
func HandleError(w http.ResponseWriter, err error) {
	log.Println(err)
	var data = map[string]string{
		"message": err.Error(),
	}

	if errors.Is(err, exceptions.AppError) {
		Response(w, data, err.Code)
		return

	}

	Response(w, data, http.StatusInternalServerError)
}

//Response sends a http response using Content-Type application/json.
func Response(w http.ResponseWriter, data interface{}, status int) {
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
