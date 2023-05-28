package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	pb "go-grpc-test/dns"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	domainFlag := flag.String("domain", "domain.xyz", "domain name to resolve")
	flag.Parse()

	grpcAddress := "0.0.0.0:9000"

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(grpcAddress, opts)
	if err != nil {
		fmt.Printf("failed to dial: %s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	req := &pb.Request{
		Domain: *domainFlag,
	}

	client := pb.NewResolverClient(conn)
	if resp, err := client.Resolve(context.Background(), req); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	} else if resp.Status == pb.Response_UNRESOLVED {
		fmt.Printf("failed to resolve domain: %s\n", *domainFlag)
		os.Exit(1)
	} else {
		fmt.Println(resp.Ip)
	}
}
