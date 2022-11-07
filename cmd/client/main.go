package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/manh737/go-grpc/protos"
	"github.com/manh737/go-grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	// serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func newLaptop(client protos.LaptopServiceClient) {
	req := &protos.CreateLaptopRequest{
		Laptop: sample.NewLaptop(),
	}
	ctx, cancer := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancer()
	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Fatalf("status code: %v, error detail: %v", st.Code(), st.Message())
		} else {

			log.Fatalf("cannot create laptop: %v", err)
		}
	}
	log.Printf("created laptop with id: %s", res.Laptop.Id)
}

// func findLaptop(client protos.LaptopServiceClient, laptopId string) bool {
// 	req := &protos.CreateLaptopRequest{
// 		Laptop: sample.NewLaptop(),
// 	}
// 	res, err := client.CreateLaptop(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("cannot create laptop: %v", err)
// 	}
// 	log.Printf("created laptop with id: %s", res.Laptop.Cpu)
// 	return req.Laptop
// }

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	// if *tls {
	// 	if *caFile == "" {
	// *caFile = data.Path("x509/ca_cert.pem")
	// 	}
	// 	creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
	// 	if err != nil {
	// 		log.Fatalf("Failed to create TLS credentials %v", err)
	// 	}
	// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// } else {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// }

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := protos.NewLaptopServiceClient(conn)

	// Create a new laptop
	newLaptop(client)

	// findLaptop(client, laptop.Id)
}
