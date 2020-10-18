//Package netowrk handles request to create_network ((to connect to wifi) - it runs validation and service.
package network

import "fmt"

//Service to create networkss
type Service struct {
}

func (s *Service) Create(n *NetworkDetails) {
	fmt.Println(*n)
	return
}
