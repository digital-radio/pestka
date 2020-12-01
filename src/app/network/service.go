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
	DbusFactory   dbusclient.DbusFactory
}

//NewService allows to create a new Service struct outside of package app.
func NewService(container *container.Container) Service {
	return Service{InterfaceName: container.InterfaceName, Scan: container.Scan, DbusFactory: container.DbusFactory}
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

	connection := s.DbusFactory.CreateDbusConnection()
	message := marshallJSON(details)
	// call := connection.Call("pl.digital_radio.Notify", 0, "c¼h", uint32(0), "", "Hallo Chaostreff!", "Ich begrüße euch herzlich zu meiner c¼h!", []string{}, map[string]dbus.Variant{}, int32(1000))

	call := connection.Call("pl.digital_radio.Notify", 0, message)
	if call.Err != nil {
		panic(call.Err)
	}
	fmt.Printf("====  CALL ====\n %s \n", call.Body)

	result := true
	if result == true {
		return nil
	}

	return errors.New("failed to connect")
}

//Get finds and lists networks in the neighbourhood
func (s *Service) Get() ([]wlist.Cell, error) {
	cells, err := s.Scan(s.InterfaceName)
	return cells, err
}
