//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"encoding/json"
	"errors"
	"fmt"

	wlist "github.com/MonkeyBuisness/golang-iwlist"
	dbusclient "github.com/digital-radio/pestka/src/dbus_client"
	"github.com/digital-radio/pestka/src/validation"

	"github.com/digital-radio/pestka/src/container"
)

//Service to create network uses container for other dependencies.
type Service struct {
	interfaceName string
	scan          container.Scan
	busFactory    dbusclient.BusFactoryInterface
	validator     *validation.Validator
}

//WifiOnResponse keeps info returned by corresponding dbus method
type WifiOnResponse struct {
	Code string `json:"code" validate:"required"`
}

//NewService allows to create a new Service struct outside of package app.
func NewService(container *container.Container, validator *validation.Validator) Service {
	return Service{
		interfaceName: container.InterfaceName,
		scan:          container.Scan,
		busFactory:    container.BusFactory,
		validator:     validator,
	}
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

	busObject := s.busFactory.CreateBusObject()

	message, err := marshallJSON(details)
	if err != nil {
		return fmt.Errorf("failed to connect to wifi - marshalling message: %w", err)
	}

	responseMessage, err := busObject.Call("pl.digitalradio.wifi_on", message)
	if err != nil {
		return fmt.Errorf("failed to connect to wifi - calling via dbus: %w", err)
	}

	response := WifiOnResponse{}
	s.validator.ParseAndValidateJSON([]byte(responseMessage), &response)

	if response.Code == "OK" {
		return nil
	}

	return errors.New("failed to connect to wifi")
}

//Get finds and lists networks in the neighbourhood
func (s *Service) Get() ([]wlist.Cell, error) {
	cells, err := s.scan(s.interfaceName)
	return cells, err
}
