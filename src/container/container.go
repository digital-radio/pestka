//Package container implements container that contains dependencies
package container

import (
	wlist "github.com/MonkeyBuisness/golang-iwlist"
	dbusclient "github.com/digital-radio/pestka/src/dbus_client"
)

//Scan returns a list of available wireless netowrks.
type Scan func(interfaceName string) ([]wlist.Cell, error)

//Container contains dependencies.
type Container struct {
	InterfaceName string
	Scan          Scan
	DbusFactory   dbusclient.BusFactoryInterface
}

//New allows to create a new Container struct outside of package container.
func New() Container {
	return Container{"wlan0", wlist.Scan, dbusclient.BusFactory{}}
}

//SetScan allows to override Scan.
func (c *Container) SetScan(scan Scan) {
	c.Scan = scan
}

//SetDbusFactory allows to override DbusFactory.
func (c *Container) SetDbusFactory(dbusFactory dbusclient.BusFactoryInterface) {
	c.DbusFactory = dbusFactory
}
