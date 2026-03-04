package main

import (
	"os"

	"github.com/AusterDev/nybl/internal/log"
	"github.com/AusterDev/nybl/services/auth"
)

func main() {
	log.InitLogger(true)
	defer log.Sync()

	log.Core("cmd/auth").Info("About to start auth server...")

	if addr, exists := os.LookupEnv("AUTH_SERVICE_LISTEN_ADDR"); exists {
		auth.RunService(addr)
	} else {
		log.Core("cmd/auth").Error("AUTH_SERVICE_LISTEN_ADDR has not been set, please set the env values accordingly.")
	}
}
