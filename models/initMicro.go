package models

import (
	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

var MicroClient client.Client

func init() {
	// Create service
	grpcClient := grpc.NewClient()
	consulRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// Create service
	srv := micro.NewService(
		micro.Registry(consulRegistry),
		micro.Client(grpcClient),
	)
	srv.Init()
	MicroClient = srv.Client()
}
