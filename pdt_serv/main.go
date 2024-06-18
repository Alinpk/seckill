package main

import (
	"user_serv/conf"
	"user_serv/controller"
	dc "user_serv/db"
	"user_serv/proto"

	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/server/grpc/v3"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

func ErrWapper(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// init configuration
	conf.InitEnv()
	ErrWapper(conf.InitConfig())
	// mail test
	ErrWapper(dc.InitDataCenter())

	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("go.micro.service.pdt_serv"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
	)

	service.Init()
	proto.RegisterCustomerHandler(service.Server(), new(controller.CustomerHandler))
	ErrWapper(service.Run())
}
