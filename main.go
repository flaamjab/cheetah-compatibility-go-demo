package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	compat "github.com/flaamjab/compatibility-minimal/compatibility_grpc"
)

var (
	addr = "stage1.kube.syncario.com:443"
)

func main() {
	config := tls.Config{}
	creds := grpc.WithTransportCredentials(credentials.NewTLS(&config))
	conn, err := grpc.Dial(addr, creds)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to %s: %s\n", addr, err)
		os.Exit(1)
	}
	defer conn.Close()

	client := compat.NewCompatibilityCheckerClient(conn)
	request := compat.CheckCompatibilityRequest{Version: "0.0.11"}
	result, err := client.CheckCompatibility(context.Background(), &request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to check compatibility:", err)
		os.Exit(1)
	}

	fmt.Println(result.Status)
}
