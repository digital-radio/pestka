package main

import (
	"fmt"

	dbus "github.com/guelfey/go.dbus"
)

type Server struct {
	id uint32
}

func (s Server) Notify(appName string, replacesId uint32, appIcon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expireTimeout int32) (ret uint32, err *dbus.Error) {
	fmt.Printf("Got Notification from %s:\n", appName)
	fmt.Printf("==== %s ====\n", summary)
	fmt.Println(body)
	fmt.Printf("==== END %s ====\n", summary)
	s.id++
	return s.id, nil
}

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	reply, err := conn.RequestName("pl.digital_radio", dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		panic("Name already taken")
	}

	s := Server{id: 0}

	conn.Export(s, "/malina", "pl.digital_radio")
	select {}
}
