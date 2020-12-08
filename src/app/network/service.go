//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

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

func marshallJSON(input interface{}) string {
	bytes, err := json.Marshal(input)

	if err != nil {
		log.Println(err)
		return ""
	}

	return string(bytes)
}

//Create connects to network specified in details
func (s *Service) Create(details *Details) error {
	fmt.Println(*details)

	busObject := s.BusFactory.CreateBusObject()

	message := marshallJSON(details)

	responseBody, err := busObject.Call("pl.digital_radio.Notify", message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("responseBody: " + responseBody)

	result := true
	if result == true {
		return nil
	}

	return errors.New("failed to connect to wifi")
}

//Get finds and lists networks in the neighbourhood
func (s *Service) Get() ([]wlist.Cell, error) {
	cells, err := s.Scan(s.InterfaceName)
	return cells, err
}
