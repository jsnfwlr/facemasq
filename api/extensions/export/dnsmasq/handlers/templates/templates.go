package templates

import (
	_ "embed"
)

//go:embed *.tmpl

func GetEmbeddedTemplates() (templates []string) {
	return []string{"dhcp.tmpl", "dns.tmpl", "header.tmpl"}
}
