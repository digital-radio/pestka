//Package service business logic of endpoint used to create network (to connect to wifi).
package service

import (
	"fmt"

	"github.com/digital-radio/pestka/src/network/handler"
)

type service struct {
}

//New allows to create a new Service struct outside of package app.
func New() service {
	return service{}
}

func (s *service) Create(n handler.NetworkDetails) {
	fmt.Println(n)
}
