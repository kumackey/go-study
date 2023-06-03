package main

import (
	"context"
	"fmt"
	"time"

	pb "helloworld/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	// Create service

	srv := micro.NewService()

	srv.Init()

	// Create client
	c := pb.NewHelloworldService(service, srv.Client())

	for {
		// Call service
		stream, err := c.ServerStream(context.Background(), &pb.ServerStreamRequest{Count: 3})
		if err != nil {
			logger.Fatal(err)
		}
		defer stream.Close()

		for i := 0; i < 3; i++ {
			var value map[string]int
			err = stream.RecvMsg(&value)
			if err != nil {
				break
			}
			fmt.Println(value["count"])
		}

		time.Sleep(1 * time.Second)
	}
}
