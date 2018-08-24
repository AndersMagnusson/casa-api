package main

import (
	"casa-api/grpc/casa/pkg/client"
	"context"
	"fmt"
)

func main() {
	c := client.Client("127.0.0.1:6262", "password", "hubba")
	ctx := context.Background()
	err := c.Run(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
