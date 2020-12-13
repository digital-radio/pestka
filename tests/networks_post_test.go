package tests_test

import (
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

func (bm *busConnectionMock) Call(method string, message string) (string, error) {
	// Record that the method was called and was passed in the value
	var args mock.Arguments = bm.Called(method, message)

	// Return values as expected. Cast them into the proper type (args.Get() returns simply interface{}).
	return args.String(0), args.Error(1)
}

type busFactoryMock struct {
	mock.Mock
	bcm *busConnectionMock
}

func (bfm *busFactoryMock) CreateBusObject() dbusclient.BusObjectInterface {
	return bfm.bcm
}

func TestPostNetworks(t *testing.T) {
	// Request is accepted and completed - wifi is turned on via dbus.

	// given
	bodyReader := strings.NewReader(`{"Ssid": "test_ssid", "Password": "test_password"}`)

	bcm := new(busConnectionMock)
	bcm.On("Call", "pl.digital_radio.Notify", "{\"Ssid\":\"test_ssid\",\"Password\":\"test_password\"}").Return("OK", nil)
	bfm := busFactoryMock{bcm: bcm}

	app := CreateTestApp()
	app.Container.SetBusFactory(&bfm)

	// when
	w := app.MakeRequest("POST", "/networks", bodyReader)

	// then
	a := assert.New(t)
	a.Equal(200, w.Code)
	a.Equal(
		`{"message":"OK"}`,
		w.Body.String(),
	)

	// Assert method "bcm.Call" was called.
	bcm.AssertExpectations(t)
}

func TestPostNetworksMissingPassword(t *testing.T) {
	// Request is validated and if 'Password' is missing, it does not try to turn on wifi via dbus.

	// given
	bodyReader := strings.NewReader(`{"Ssid": "test_ssid"}`)

	bcm := new(busConnectionMock)
	bfm := busFactoryMock{bcm: bcm}

	app := CreateTestApp()
	app.Container.SetBusFactory(&bfm)

	// when
	w := app.MakeRequest("POST", "/networks", bodyReader)

	// then
	a := assert.New(t)
	a.Equal(400, w.Code)
	a.Equal(
		`{"message":"Bad request: invalid input - Key: 'createNetworkRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		w.Body.String(),
	)

	// Assert method "bcm.Call" was not called.
	bcm.AssertNotCalled(t, "Call")
}

func TestPostNetworksMissingSsid(t *testing.T) {
	// Request is validated and if 'Ssid' is missing, it does not try to turn on wifi via dbus.

	// given
	bodyReader := strings.NewReader(`{"Password": "test_password"}`)

	bcm := new(busConnectionMock)
	bfm := busFactoryMock{bcm: bcm}

	app := CreateTestApp()
	app.Container.SetBusFactory(&bfm)

	// when
	w := app.MakeRequest("POST", "/networks", bodyReader)

	// then
	a := assert.New(t)
	a.Equal(400, w.Code)
	a.Equal(
		`{"message":"Bad request: invalid input - Key: 'createNetworkRequest.Ssid' Error:Field validation for 'Ssid' failed on the 'required' tag"}`,
		w.Body.String(),
	)

	// Assert method "bcm.Call" was not called.
	bcm.AssertNotCalled(t, "Call")
}
