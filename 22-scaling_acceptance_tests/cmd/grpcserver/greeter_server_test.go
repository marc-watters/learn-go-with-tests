package main_test

import (
	"fmt"
	"go-specs-greet/v2/adapters"
	"go-specs-greet/v2/adapters/grpcserver"
	"go-specs-greet/v2/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
}
