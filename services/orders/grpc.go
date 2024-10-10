package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()

	// register gRPC services

	log.Println("Starting gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}
