package tests

import (
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"

	"github.com/digital-radio/pestka/src/app"
	"github.com/digital-radio/pestka/src/container"
)

//TestApp is a struct simulating app during tests.
type TestApp struct {
	Container container.Container
}

//CreateTestApp allows to create an app just for testing.
func CreateTestApp() TestApp {
	return TestApp{container.New()}
}

//MakeRequest allows to simulate http server during tests.
func (a *TestApp) MakeRequest(method, target string, body io.Reader) *httptest.ResponseRecorder {
	var application = app.New(a.Container)
	var router = application.CreateRouter()
	var w = httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest(method, target, body))
	return w
}

//MarshallJson helper method to marshall json.
func MarshallJson(input interface{}) string {
	bytes, err := json.Marshal(input)

	if err != nil {
		log.Println(err)
		return ""
	}

	return string(bytes)
}

//JsonRemarshal helper method to re-marshall json.
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
