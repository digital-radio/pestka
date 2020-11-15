package tests

import (
	"encoding/json"
	"github.com/digital-radio/pestka/src/app"
	"github.com/digital-radio/pestka/src/container"
	"io"
	"log"
	"net/http/httptest"
)

type TestApp struct {
	Container container.Container
}

func (a *TestApp) MakeRequest(method, target string, body io.Reader) *httptest.ResponseRecorder {
	var application = app.New(a.Container)
	var router = application.CreateRouter()
	//FIXME Co to robi ?
	var w = httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest(method, target, body))
	return w
}

func MarshallJson(input interface{}) string {
	bytes, err := json.Marshal(input)

	if err != nil {
		log.Println(err)
		return ""
	}

	return string(bytes)
}

func JsonRemarshal(bytes []byte) []byte {
	var ifce interface{}
	err := json.Unmarshal(bytes, &ifce)
	if err != nil {
		return []byte{}
	}
	output, err := json.Marshal(ifce)
	if err != nil {
		return []byte{}
	}
	return output
}
