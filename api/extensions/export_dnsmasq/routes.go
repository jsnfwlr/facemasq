package main

import (
	"facemasq/extensions/export_dnsmasq/handlers"
	"facemasq/lib/extensions"
)

func getRoutes() []extensions.RouteDefinition {
	return []extensions.RouteDefinition{
		{Path: `/export/dnsmasq/dhcp`, Handler: handlers.WriteDHCPConfig, Methods: "GET", Name: "WriteDHCPConfigGet"},
		{Path: `/export/dnsmasq/dhcp`, Handler: handlers.WriteDHCPConfig, Methods: "PUT", Name: "WriteDHCPConfigPut"},
		{Path: `/export/dnsmasq/dns`, Handler: handlers.WriteDNSConfig, Methods: "GET", Name: "WriteDNSConfigGet"},
		{Path: `/export/dnsmasq/dns`, Handler: handlers.WriteDNSConfig, Methods: "PUT", Name: "WriteDNSConfigPut"},
	}
}
