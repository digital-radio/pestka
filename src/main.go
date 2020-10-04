//Package main package implements creation and start of http server.
package main

import (
	"fmt"
	"github.com/digital-radio/pestka/src/app"
	"github.com/digital-radio/pestka/src/container"
	"log"
	"net/http"
)

func main() {
	var c = container.New()
	var a = app.New(c)
	var router = a.CreateRouter()
	var port = 8080
	var addr = fmt.Sprint(":", port)

	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
