package auth

import (
	"github.com/AusterDev/nybl/internal/httpserver"
	"github.com/AusterDev/nybl/internal/log"
	"github.com/gofiber/fiber/v2"
)

func RunAuthService(listenAddr string) {
	log.Service("auth").Info("Running auth service...")
	httpserver.StartHttpServer(listenAddr, fiber.Config{})
}