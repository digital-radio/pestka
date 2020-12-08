package tests_test

import (
	"fmt"
	"strings"
	"testing"

	dbusclient "github.com/digital-radio/pestka/src/dbus_client"
	. "github.com/digital-radio/pestka/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type busConnectionMock struct {
	mock.Mock
}

func (busConnectionMock) Call(method string, message string) (string, error) {
	return "a", nil
}

type busFactoryMock struct {
	mock.Mock
}

func (busFactoryMock) CreateBusObject() dbusclient.BusObjectInterface {
	bcm := new(busConnectionMock)
	fmt.Printf("Create bus object !!! \n")

	bcm.On("Call", "{\"ssid\": \"test_ssid\", \"password\": \"test_password\"}").Return(("b"))
	return bcm

}

func TestPostNetworks(t *testing.T) {
	// given
	bodyReader := strings.NewReader(`{"ssid": "test_ssid", "password": "test_password"}`)

	bfm := new(busFactoryMock)

	app := CreateTestApp()

	app.Container.SetBusFactory(bfm)

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

func TestPostNetworksMissingPassword(t *testing.T) {
	// given

	bodyReader := strings.NewReader(`{"ssid": "test_ssid"}`)

	app := TestApp{}

	// when
	w := app.MakeRequest("POST", "/networks", bodyReader)

	// then
	a := assert.New(t)
	a.Equal(400, w.Code)
}

func TestPostNetworksMissingSsid(t *testing.T) {
	// given

	bodyReader := strings.NewReader(`{"password": "test_password"}`)

	app := TestApp{}

	// when
	w := app.MakeRequest("POST", "/networks", bodyReader)

	// then
	a := assert.New(t)
	a.Equal(400, w.Code)
}
