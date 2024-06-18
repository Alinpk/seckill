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
	// ErrWapper(utils.SendVerifyMail([]byte("18810535172@163.com"), []byte("2133")))
	ErrWapper(dc.InitDataCenter())

	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("go.micro.service.user_serv"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
		// micro.Client(grpc.NewClient()),
	)

	service.Init()
	proto.RegisterCustomerHandler(service.Server(), new(controller.CustomerHandler))
	proto.RegisterSellerHandler(service.Server(), new(controller.SellerHandler))
	ErrWapper(service.Run())
}
