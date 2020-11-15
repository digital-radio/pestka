package tests_test

import (
	"strings"
	"testing"

	"github.com/digital-radio/pestka/src/container"
	. "github.com/digital-radio/pestka/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedScan struct {
	mock.Mock
	container.Scan
}

func TestPostNetworksList(t *testing.T) {
	// given

	bodyReader := strings.NewReader(`{"ssid": "test_ssid", "password": "test_password"}`)

	app := TestApp{}

	// when
	w := app.MakeRequest("POST", "/networks", bodyReader)

	// then
	a := assert.New(t)
	a.Equal(200, w.Code)
	a.Equal(
		`{"message":"OK"}`,
		w.Body.String(),
	)
}
