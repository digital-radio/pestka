//Package dbus implements handling of the dbus
package dbus

import (
	"github.com/godbus/dbus"
)

type DbusFactory struct{}


func (*DbusFactory) CreateDbusConnection() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	// func (conn *Conn) Object(dest string, path ObjectPath) *Object
	obj := conn.Object("pl.digital_radio", "/malina")

	return obj
}
