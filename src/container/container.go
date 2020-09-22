package container

import wlist "github.com/MonkeyBuisness/golang-iwlist"

type Scan = func (interfaceName string) ([]wlist.Cell, error)

type Container struct {
	InterfaceName string
	Scan          Scan
}

func New() Container {
	return Container{"wlan0", wlist.Scan}
}

func (c *Container) SetScan(scan Scan) {
	c.Scan = scan
}
