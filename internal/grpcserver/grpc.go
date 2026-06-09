package grpcserver

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger,addr string) *GRPCServer {
	return &GRPCServer{addr: addr}
}

func (s *GRPCServer) StartGRPCServer() {
	tcp, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.logger.Errorw("Failed to create TCP instance", "err", err, "addr", s.addr)
	}

	grpcServer := grpc.NewServer()
	if err = grpcServer.Serve(tcp); err != nil {
		s.logger.Errorw("Failed to create GRPC server instance", "err", err, "addr", s.addr)
	}
}