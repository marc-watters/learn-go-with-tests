package main_test

import (
	"fmt"
	"go-specs-greet/v2/adapters"
	"go-specs-greet/v2/adapters/grpcserver"
	"go-specs-greet/v2/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
