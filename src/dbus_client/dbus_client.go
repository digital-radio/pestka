//Package dbus_client implements sending messages via dbus
package dbusclient

import (
	"github.com/godbus/dbus"
)

type DbusFactory struct{}

func (*DbusFactory) CreateDbusConnection() dbus.BusObject {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	// func (conn *Conn) Object(dest string, path ObjectPath) *Object
	obj := conn.Object("pl.digital_radio", "/malina")

	return obj
}
