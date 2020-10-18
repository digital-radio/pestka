//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import "fmt"

//Service to create network
type Service struct {
}

//Create connects to network specified in details
func (s *Service) Create(details *Details) {
	fmt.Println(*details)
	return
}
