package main

import (
	"user_serv/conf"
	"user_serv/controller"
	"user_serv/proto"

	"github.com/asim/go-micro/plugins/registry/etcd/v3"
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

	service := micro.NewService(
		micro.Name("go.micro.service.user_serv"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
	)

	service.Init()
	proto.RegisterCustomerHandler(service.Server(), new(controller.CustomerHandler))
	ErrWapper(service.Run())
}
