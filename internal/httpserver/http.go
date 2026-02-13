package httpserver

import (
	"github.com/AusterDev/nybl/internal/log"
	"github.com/gofiber/fiber/v2"
)

/*
Run me in a goroutine
*/
func StartHttpServer(addr string, config fiber.Config) {
	server := fiber.New(config)

	logger := log.Core("http-start")

	err := server.Listen(addr)
	if err != nil {
		logger.Errorw("Failed to start http server", "err", err)
	}
}