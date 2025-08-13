package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/dimahum/starlink-cntrl-grpc/proto"
)

var (
	addr     = flag.String("addr", "192.168.100.1:9200", "Starlink device address")
	insecureFlag = flag.Bool("insecure", true, "Use insecure connection")
	timeout  = flag.Duration("timeout", 10*time.Second, "Connection timeout")
)

func main() {
	flag.Parse()

	fmt.Printf("Connecting to Starlink device at %s...\n", *addr)

	// Create connection options
	var opts []grpc.DialOption
	if *insecureFlag {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// For TLS connections, accept any certificate
		config := &tls.Config{
			InsecureSkipVerify: true,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	}

	// Set up a connection to the server
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create the client
	client := pb.NewDeviceClient(conn)

	// Get status
	err = getStatus(client)
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}
}

func getStatus(client pb.DeviceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create the request
	req := &pb.Request{
		Request: &pb.Request_GetStatus{
			GetStatus: &pb.GetStatusRequest{},
		},
	}

	fmt.Println("Sending GetStatus request...")

	// Make the request
	resp, err := client.Handle(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call Handle: %v", err)
	}

	// Parse the response
	switch r := resp.Response.(type) {
	case *pb.Response_GetStatus:
		status := r.GetStatus
		fmt.Println("\n=== Starlink Status ===")
		
		if status.DeviceInfo != nil {
			fmt.Printf("Device ID: %s\n", status.DeviceInfo.Id)
			fmt.Printf("Hardware Version: %s\n", status.DeviceInfo.HardwareVersion)
			fmt.Printf("Software Version: %s\n", status.DeviceInfo.SoftwareVersion)
			fmt.Printf("Country Code: %s\n", status.DeviceInfo.CountryCode)
		}
		
		if status.DeviceState != nil {
			fmt.Printf("Uptime: %.2f seconds\n", status.DeviceState.UptimeS)
		}
		
		fmt.Println("========================")
		
	default:
		return fmt.Errorf("unexpected response type: %T", resp.Response)
	}

	return nil
}