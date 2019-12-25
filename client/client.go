package main

import (
	"context"
	"log"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	r, err := c.Echo(context.Background(), &pb.EchoRequest{Message: "This is from client..."})
	if err != nil {
		log.Fatalf("Fail to call Echo：%v", err)
	}
	log.Printf("Result：%s , Time: %d", r.Message, r.Timestamp)
}
