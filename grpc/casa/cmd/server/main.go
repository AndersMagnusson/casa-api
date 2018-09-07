package main

import (
	"casa-api/grpc/casa/pkg/server"
	"context"
	"fmt"
)

func main() {
	s := server.Server(":50051", "password")
	ctx := context.Background()
	err := s.Run(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
