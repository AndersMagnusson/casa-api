package main

import (
	"casa-api/grpc/casa/pkg/server"
	"context"
	"fmt"
)

func main() {
	s := server.Server("127.0.0.1:6262", "password")
	ctx := context.Background()
	err := s.Run(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
