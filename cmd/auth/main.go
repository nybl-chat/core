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
	auth.RunAuthService(os.Getenv("AUTH_SERVICE_LISTEN_ADDR"))
}