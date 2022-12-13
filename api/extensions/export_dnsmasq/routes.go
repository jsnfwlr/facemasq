package main

import (
	"facemasq/extensions/export_dnsmasq/handlers"

	"github.com/uptrace/bunrouter"
)

func getRoutes(router *bunrouter.Router) {
	router.WithGroup("/export/dnsmasq", func(group *bunrouter.Group) {
		group.GET("/dhcp", handlers.WriteDHCPConfig) //"WriteDHCPConfigGet"
		group.PUT("/dhcp", handlers.WriteDHCPConfig) // "WriteDHCPConfigPut"
		group.GET("/dns", handlers.WriteDNSConfig)   // "WriteDNSConfigGet"
		group.PUT("/dns", handlers.WriteDNSConfig)   // "WriteDNSConfigPut"
	})
}
