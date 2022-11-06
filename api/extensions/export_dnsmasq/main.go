package main

import (
	"facemasq/lib/events"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
)

func LoadExtension(manager *extensions.Manager) (err error) {
	logging.Debug1("Loading Extension")
	manager.RegisterRoutes(getRoutes())
	manager.RegisterListeners([]extensions.Listener{
		{Kind: `device:after:.*`, Listener: ExportOnSave},
		{Kind: `plugin:loaded:dnsmasq.so`, Listener: AlertOnLoad},
	})
	return
}

func ExportOnSave(e events.Event) {
	logging.System("dnsmasq ExportOnSave %+v ", e)
}

func AlertOnLoad(e events.Event) {
	logging.System("dnsmasq AlertOnLoad %+v ", e)
}
