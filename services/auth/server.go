package auth

import (
	"github.com/AusterDev/nybl/internal/grpcserver"
	"github.com/AusterDev/nybl/internal/log"
)

func RunService(addr string) {
	logger := log.Service("auth")
	logger.Info("Running auth service...")

	s := grpcserver.New(logger, addr)
	s.StartGRPCServer()
}
