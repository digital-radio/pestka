package tests_test

import (
	wlist "github.com/MonkeyBuisness/golang-iwlist"
	"github.com/digital-radio/pestka/src/container"
	. "github.com/digital-radio/pestka/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedScan struct {
	mock.Mock
	container.Scan
}

func TestGetNetworksList(t *testing.T) {
	// given
	mockedScan := new(MockedScan)
	mockedScan.On("scan")

	app := TestApp{}
	app.Container.SetScan(func(interfaceName string) ([]wlist.Cell, error) {
		return []wlist.Cell {
			{
				CellNumber: "example",
				MAC:        "example_mac",
			},
		}, nil
	})

	// when
	w := app.MakeRequest("GET", "/networks", nil)

	// then
	a := assert.New(t)
	a.Equal(200, w.Code)
	a.Equal(
		MarshallJson([]map[string]interface{}{
			{
				"cell_number": "example",
				"mac": "example_mac",
				"essid": "",
				"mode": "",
				"frequency": 0,
				"frequency_units": "",
				"channel": 0,
				"encryption_key": false,
				"encryption": "",
				"signal_quality": 0,
				"signal_total": 0,
				"signal_level": 0,
			},
		}),
		string(JsonRemarshal(w.Body.Bytes())),
	)
}
