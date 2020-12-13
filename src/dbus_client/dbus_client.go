//Package dbusclient implements sending messages via dbus
package dbusclient

import (
	"fmt"

	"github.com/godbus/dbus"
)

//BusObjectInterface defines available methods on BusObject.
type BusObjectInterface interface {
	Call(method string, message string) (string, error)
}

//BusFactoryInterface defines method to create bus object.
type BusFactoryInterface interface {
	CreateBusObject() BusObjectInterface
}

//BusObject gathers methods to communicate via dbus.
type BusObject struct {
	object dbus.BusObject
}

//Call sends messages via dbus.
func (bo *BusObject) Call(method string, message string) (string, error) {
	// call := connection.Call("pl.digital_radio.Notify", 0, "c¼h", uint32(0), "", "Hallo Chaostreff!", "Ich begrüße euch herzlich zu meiner c¼h!", []string{}, map[string]dbus.Variant{}, int32(1000))
	var call *dbus.Call = bo.object.Call(method, 0, message)
	if call.Err != nil {
		panic(call.Err)
		// return errors.New("failed to send message")

	}
	fmt.Printf("====  CALL ====\n %s \n", call.Body)

	return "ala", nil

}

//BusFactory has method to create bus object.
type BusFactory struct{}

//CreateBusObject creates bus object and connects to dbus.
func (*BusFactory) CreateBusObject() BusObjectInterface {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	// func (conn *Conn) Object(dest string, path ObjectPath) *Object
	var obj dbus.BusObject = conn.Object("pl.digital_radio", "/malina")

	return &BusObject{object: obj}
}
