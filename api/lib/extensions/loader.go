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

	"github.com/uptrace/bunrouter"
)

type ExtensionManager struct {
	Router      *bunrouter.Router
	listeners   []Listener
	coordinator events.Manager
	Extensions  []Extension
	numRoutes   int
}

type Extension struct {
	Name     string
	Filename string
}

var Manager *ExtensionManager

type Loader interface {
	LoadExtension(*ExtensionManager) (extensionName string, err error)
}

// LoadExtensions loads extensions from the directory with the given path, looking for
// all .so files in there. It creates a new PluginManager and registers the
// plugins with it.
func LoadExtensions(router *bunrouter.Router) (manager *ExtensionManager, err error) {
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

	Manager = newManager(router)
	for f := range files {
		if !files[f].IsDir() && strings.HasSuffix(files[f].Name(), ".so") {
			fullpath := filepath.Join(path, files[f].Name())

			extension, err = plugin.Open(fullpath)
			if err != nil {
				logging.Error("Could not inspect %s: %+v", files[f].Name(), err)
				err = nil
				continue
			}

			symLoader, err = extension.Lookup("LoadExtension")
			if err != nil {
				logging.Error("Could not find LoadExtension for %s: %+v", files[f].Name(), err)
				err = nil
				continue
			}

			loader := symLoader.(func(*ExtensionManager) (string, error))
			extensionName, err = loader(Manager)
			if err != nil {
				logging.Error("Could not load %s: %+v", files[f].Name(), err)
				err = nil
				continue
			}
			logging.Info("Loaded extension: %s", extensionName)

			Manager.Extensions = append(Manager.Extensions, Extension{Name: extensionName, Filename: files[f].Name()})
		}
	}
	logging.Info("%d listeners registered | %d routes registered", len(Manager.listeners), 0)
	for l := range Manager.listeners {
		Manager.coordinator.Listen(Manager.listeners[l].Kind, Manager.listeners[l].Listener)
	}

	for p := range Manager.Extensions {
		err = Manager.coordinator.Emit(events.Event{Kind: fmt.Sprintf("plugin:loaded:%s", Manager.Extensions[p])})
		if err != nil {
			logging.Error("Error with event: %v", err)
		}
	}
	manager = Manager
	return
}

func newManager(router *bunrouter.Router) *ExtensionManager {
	return &ExtensionManager{
		coordinator: events.DefaultManager(),
		Router:      router,
	}
}

// func (manager *ExtensionManager) RegisterRoutes(routes []RouteDefinition) {
// 	manager.numRoutes = manager.numRoutes + len(routes)
// 	manager.routes = append(manager.routes, routes...)
// }

func (manager *ExtensionManager) RegisterListeners(listeners []Listener) {
	logging.Debug("Registering event listeners")
	for l := range listeners {
		logging.Info("%s", listeners[l].Kind)
	}
	manager.listeners = append(manager.listeners, listeners...)
}

func (manager *ExtensionManager) GetRoutes(router *bunrouter.Router) {

}

func (manager *ExtensionManager) HasRoutes() bool {
	return (manager.numRoutes > 0)
}

func (manager *ExtensionManager) GetListeners() (listeners []Listener) {
	return manager.listeners
}

func (manager *ExtensionManager) GetCoordinator() (coordinator events.Manager) {
	return manager.coordinator
}
