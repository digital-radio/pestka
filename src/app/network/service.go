//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"fmt"

	wlist "github.com/MonkeyBuisness/golang-iwlist"
	"github.com/godbus/dbus"

	"github.com/digital-radio/pestka/src/container"
)

//Service to create network uses container for other dependencies.
type Service struct {
	Scan        container.Scan
	DbusFactory dbus.DbusFactory
}

//NewService allows to create a new Service struct outside of package app.
func NewService(container *container.Container) Service {
	return Service{Scan: container.Scan, DbusFactory: container.DbusFactory}
}

//Create connects to network specified in details
func (s *Service) Create(details *Details) {
	fmt.Println(*details)

	connection := s.DbusFactory.CreateDbusConnection()
	call := connection.Call("pl.digital_radio.Notify", 0, "c¼h", uint32(0), "", "Hallo Chaostreff!", "Ich begrüße euch herzlich zu meiner c¼h!", []string{}, map[string]dbus.Variant{}, int32(1000))
	if call.Err != nil {
		panic(call.Err)
	}
	fmt.Printf("====  CALL ====\n %s \n", call.Body)
}

//Get finds and lists networks in the neighbourhood
func (s *Service) Get() ([]wlist.Cell, error) {
	cells, err := s.Scan(s.container.InterfaceName)
	return cells, err
}
