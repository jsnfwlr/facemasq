package main

import (
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/pkg/errors"
)

const ExtensionName = "Omada Importer"

var (
	frequency time.Duration
	mongoDSN  string
)

func init() {
	var err error
	frequency, err = time.ParseDuration(os.Getenv("OMADA_FREQUENCY"))
	if err != nil {
		frequency = time.Duration(60) * time.Second
	}
	mongoDSN = os.Getenv("OMADA_DSN")

}

func LoadExtension(manager *extensions.ExtensionManager) (extensionName string, err error) {
	extensionName = ExtensionName
	if mongoDSN == "" {
		err = errors.New("the Omada importer requires the OMADA_DSN environment variable be set")
		return
	}
	scheduleImport()
	/*
		manager.RegisterRoutes(getRoutes())
		manager.RegisterListeners([]extensions.Listener{
			{Kind: `device:after:.*`, Listener: ExportOnSave},
			{Kind: `plugin:loaded:dnsmasq.so`, Listener: AlertOnLoad},
		})
	*/
	return
}

func scheduleImport() {
	sched := gocron.NewScheduler(time.UTC)
	sched.Every(frequency).Do(func() {
		logging.Info("Doing Omada import")
	})
	sched.StartAsync()
}
