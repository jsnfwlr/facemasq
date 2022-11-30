package extensions

import (
	"facemasq/lib/files"
	"facemasq/lib/logging"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"facemasq/lib/events"
)

type ExtensionManager struct {
	routes      []RouteDefinition
	listeners   []Listener
	coordinator events.Manager
	Extensions  []Extension
}

type Extension struct {
	Name     string
	Filename string
}

var Manager *ExtensionManager

type Loader interface {
	LoadExtension(*ExtensionManager) (extensionName string, err error)
}

// LoadPlugins loads plugins from the directory with the given path, looking for
// all .so files in there. It creates a new PluginManager and registers the
// plugins with it.
func LoadPlugins() (Manager *ExtensionManager, err error) {
	path, err := files.GetDir("extensions")
	if err != nil {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	var extension *plugin.Plugin
	var symLoader plugin.Symbol
	var extensionName string

	Manager = newManager()
	for f := range files {
		if !files[f].IsDir() && strings.HasSuffix(files[f].Name(), ".so") {
			fullpath := filepath.Join(path, files[f].Name())

			extension, err = plugin.Open(fullpath)
			if err != nil {
				logging.Error("Could not load %s: %+v", files[f].Name(), err)
				continue
			}

			symLoader, err = extension.Lookup("LoadExtension")
			if err != nil {
				logging.Error("Could not find LoadExtension for %s: %+v", files[f].Name(), err)
				continue
			}

			loader := symLoader.(func(*ExtensionManager) (string, error))
			extensionName, err = loader(Manager)
			if err != nil {
				logging.Error("Could not load %s: %+v", files[f].Name(), err)
				continue
			}
			logging.Info("Loaded extension: %s", extensionName)

			Manager.Extensions = append(Manager.Extensions, Extension{Name: extensionName, Filename: files[f].Name()})
		}
	}

	logging.Info("%d listeners registered", len(Manager.listeners))
	for l := range Manager.listeners {
		Manager.coordinator.Listen(Manager.listeners[l].Kind, Manager.listeners[l].Listener)
	}

	for p := range Manager.Extensions {
		err = Manager.coordinator.Emit(events.Event{Kind: fmt.Sprintf("plugin:loaded:%s", Manager.Extensions[p])})
		if err != nil {
			logging.Error("Error with event: %v", err)
		}
	}
	return
}

func newManager() *ExtensionManager {
	return &ExtensionManager{
		coordinator: events.DefaultManager(),
	}
}

func (manager *ExtensionManager) RegisterRoutes(routes []RouteDefinition) {
	manager.routes = append(manager.routes, routes...)
}

func (manager *ExtensionManager) RegisterListeners(listeners []Listener) {
	logging.Debug("Registering event listeners")
	for l := range listeners {
		logging.Info("%s", listeners[l].Kind)
	}
	manager.listeners = append(manager.listeners, listeners...)
}

func (manager *ExtensionManager) GetRoutes() (routes []RouteDefinition) {
	return manager.routes
}

func (manager *ExtensionManager) GetListeners() (listeners []Listener) {
	return manager.listeners
}

func (manager *ExtensionManager) GetCoordinator() (coordinator events.Manager) {
	return manager.coordinator
}
