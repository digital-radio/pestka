//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"encoding/json"
	"errors"
	"fmt"

	wlist "github.com/MonkeyBuisness/golang-iwlist"
	dbusclient "github.com/digital-radio/pestka/src/dbus_client"

	"github.com/digital-radio/pestka/src/container"
)

//Service to create network uses container for other dependencies.
type Service struct {
	InterfaceName string
	Scan          container.Scan
	BusFactory    dbusclient.BusFactoryInterface
}

//NewService allows to create a new Service struct outside of package app.
func NewService(container *container.Container) Service {
	return Service{InterfaceName: container.InterfaceName, Scan: container.Scan, BusFactory: container.BusFactory}
}

func marshallJSON(input interface{}) (string, error) {
	bytes, err := json.Marshal(input)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//Create connects to network specified in details
func (s *Service) Create(details *Details) error {

	busObject := s.BusFactory.CreateBusObject()

	message, err := marshallJSON(details)
	if err != nil {
		return fmt.Errorf("failed to connect to wifi - marshalling message: %w", err)
	}

	responseBody, err := busObject.Call("pl.digital_radio.Notify", message)
	if err != nil {
		return fmt.Errorf("failed to connect to wifi - calling via dbus: %w", err)
	}

	if responseBody == "OK" {
		return nil
	}

	return errors.New("failed to connect to wifi")
}

//Get finds and lists networks in the neighbourhood
func (s *Service) Get() ([]wlist.Cell, error) {
	cells, err := s.Scan(s.InterfaceName)
	return cells, err
}
