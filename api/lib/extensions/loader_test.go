package extensions

import "testing"

func TestLoadingExtensions(t *testing.T) {
	manager, err := LoadPlugins()
	if err != nil {
		t.Fatalf("%v", err)
	}
	routes := manager.GetRoutes()
	for r := range routes {
		t.Log(routes[r].Name)
	}
}