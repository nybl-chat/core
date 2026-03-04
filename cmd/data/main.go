package main

import (
	"os"

	"github.com/AusterDev/nybl/internal/log"
	"github.com/AusterDev/nybl/services/data"
)

func main() {
	log.InitLogger(true)
	defer log.Sync()

	lmod := log.Core("cmd/data")
	lmod.Info("About to start data server...")

	if addr, exists := os.LookupEnv("DATA_SERVICE_LISTEN_ADDR"); exists {
		data.RunService(addr)
	} else {
		lmod.Error("DATA_SERVICE_LISTEN_ADDR has not been set, please set the env values accordingly.")
	}
}
