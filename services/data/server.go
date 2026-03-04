package data

import (
	"github.com/AusterDev/nybl/internal/httpserver"
	"github.com/AusterDev/nybl/internal/log"
	"github.com/gofiber/fiber/v2"
)

func RunService(addr string) {
	log.Service("auth").Info("Running data service...")
	httpserver.StartHttpServer(addr, fiber.Config{})
}
