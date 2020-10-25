//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"fmt"

	wlist "github.com/MonkeyBuisness/golang-iwlist"

	"github.com/digital-radio/pestka/src/container"
)

//Service to create network uses container for other dependencies.
type Service struct {
	container *container.Container
}

//NewService allows to create a new Service struct outside of package app.
func NewService(container *container.Container) Service {
	return Service{container}
}

//Create connects to network specified in details
func (s *Service) Create(details *Details) {
	fmt.Println(*details)
	return
}

//Get finds and lists networks in the neighbourhood
func (s *Service) Get() ([]wlist.Cell, error) {
	cells, err := s.container.Scan(s.container.InterfaceName)
	return cells, err
}
