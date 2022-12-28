package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	tgpb "tgbot/pb"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	tgpb.RegisterTGServer(grpcServer, &GRPCServer{})
	grpcServer.Serve(listener)

}
