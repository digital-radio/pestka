package main

import (
	"fmt"

	"github.com/godbus/dbus"
)

func main() {

	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	// func (conn *Conn) Object(dest string, path ObjectPath) *Object
	obj := conn.Object("pl.digital_radio", "/malina")

	// Interface from the specification:
	// UINT32 org.freedesktop.Notifications.Notify (STRING app_name, UINT32 replaces_id, STRING app_icon, STRING summary, STRING body, ARRAY actions, DICT hints, INT32 expire_timeout);

	// func (o *Object) Call(method string, flags Flags, args ...interface{}) *Call
	call := obj.Call("pl.digital_radio.Notify", 0, "c¼h", uint32(0), "", "Hallo Chaostreff!", "Ich begrüße euch herzlich zu meiner c¼h!", []string{}, map[string]dbus.Variant{}, int32(1000))
	if call.Err != nil {
		panic(call.Err)
	}
	fmt.Printf("====  CALL ====\n %s \n", call.Body)

}
