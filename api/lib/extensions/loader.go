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

type Manager struct {
	routes      []RouteDefinition
	listeners   []Listener
	coordinator events.Manager
}

var Extensions *Manager

type Loader interface {
	LoadExtension(*Manager) (err error)
}

// LoadPlugins loads plugins from the directory with the given path, looking for
// all .so files in there. It creates a new PluginManager and registers the
// plugins with it.
func LoadPlugins() (*Manager, error) {
	path, err := files.GetDir("extensions")
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var extension *plugin.Plugin
	var symLoader plugin.Symbol

	manager := newManager()
	var pluginsLoaded []string
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
				logging.Error("Could not load %s: %+v", files[f].Name(), err)
				continue
			}

			loader := symLoader.(func(*Manager) error)
			err = loader(manager)
			if err != nil {
				logging.Error("Could not load %s: %+v", files[f].Name(), err)
				continue
			}
			logging.System("Loaded extension %s", files[f].Name())
			pluginsLoaded = append(pluginsLoaded, files[f].Name())
		}
	}
	Extensions = manager

	logging.System("%d listeners registered", len(manager.listeners))
	for l := range manager.listeners {
		manager.coordinator.Listen(manager.listeners[l].Kind, manager.listeners[l].Listener)
	}

	for p := range pluginsLoaded {
		err = manager.coordinator.Emit(events.Event{Kind: fmt.Sprintf("plugin:loaded:%s", pluginsLoaded[p])})
		if err != nil {
			logging.Error("Error with event: %v", err)
		}
	}
	return manager, nil
}

func newManager() *Manager {
	manager := &Manager{}
	manager.coordinator = events.DefaultManager()
	return manager
}

func (manager *Manager) RegisterRoutes(routes []RouteDefinition) {
	manager.routes = append(manager.routes, routes...)
}

func (manager *Manager) RegisterListeners(listeners []Listener) {
	logging.Debug3("Registering event listeners")
	for l := range listeners {
		logging.System("%s", listeners[l].Kind)
	}
	manager.listeners = append(manager.listeners, listeners...)
}

func (manager *Manager) GetRoutes() (routes []RouteDefinition) {
	return manager.routes
}

func (manager *Manager) GetListeners() (listeners []Listener) {
	return manager.listeners
}

func (manager *Manager) GetCoordinator() (coordinator events.Manager) {
	return manager.coordinator
}
