package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "go-grpc-test/dns"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type dnsServer struct {
	pb.ResolverServer
}

func (s *dnsServer) Resolve(_ context.Context, req *pb.Request) (*pb.Response, error) {
	domainIpMap := map[string]string{
		"domain.xyz": "192.168.0.0",
		"test.xyz":   "192.168.0.1",
	}

	if ip, ok := domainIpMap[req.GetDomain()]; !ok {
		return &pb.Response{
			Status: pb.Response_UNRESOLVED,
			Ip:     "",
		}, status.Errorf(codes.NotFound, "domain not found [%s]", req.Domain)
	} else {
		return &pb.Response{
			Status: pb.Response_RESOLVED,
			Ip:     ip,
		}, nil
	}
}

func main() {
	grpcAddress := "0.0.0.0:9000"

	fmt.Printf("staring DNS GRPC server on %v\n", grpcAddress)

	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		fmt.Printf("failed to listen on %v, error: %s\n", grpcAddress, err.Error())
		os.Exit(1)
	}

	fmt.Printf("gRPC server started listening on %v\n", grpcAddress)

	grpcServer := grpc.NewServer()
	pb.RegisterResolverServer(grpcServer, &dnsServer{})

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
