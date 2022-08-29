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
	err := connect(addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func connect(addr string) error {
	config := tls.Config{}
	creds := grpc.WithTransportCredentials(credentials.NewTLS(&config))
	conn, err := grpc.Dial(addr, creds)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %s", addr, err)
	}
	defer conn.Close()

	client := compat.NewCompatibilityCheckerClient(conn)
	request := compat.CheckCompatibilityRequest{Version: "0.0.11"}
	result, err := client.CheckCompatibility(context.Background(), &request)
	if err != nil {
		return fmt.Errorf("failed to check compatibility: %s", err)
	}

	fmt.Println(result.Status)

	return nil
}
