//Package container implements container that contains dependencies
package container

import wlist "github.com/MonkeyBuisness/golang-iwlist"

//Scan returns a list of available wireless netowrks.
type Scan func(interfaceName string) ([]wlist.Cell, error)

//Container contains dependencies.
type Container struct {
	InterfaceName string
	Scan          Scan
}

//New allows to create a new Container struct outside of package container.
func New() Container {
	return Container{"wlan0", wlist.Scan}
}

//SetScan allows to override Scan.
func (c *Container) SetScan(scan Scan) {
	c.Scan = scan
}
