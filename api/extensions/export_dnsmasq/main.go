package main

import (
	"facemasq/lib/events"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
)

const ExtensionName = "DNSMasq Exporter"

func LoadExtension(manager *extensions.ExtensionManager) (extensionName string, err error) {
	extensionName = ExtensionName
	manager.RegisterRoutes(getRoutes())
	manager.RegisterListeners([]extensions.Listener{
		{Kind: `device:after:.*`, Listener: ExportOnSave},
		{Kind: `plugin:loaded:dnsmasq.so`, Listener: AlertOnLoad},
	})
	return
}

func ExportOnSave(e events.Event) {
	logging.Info("dnsmasq ExportOnSave %+v ", e)
}

func AlertOnLoad(e events.Event) {
	logging.Info("dnsmasq AlertOnLoad %+v ", e)
}
