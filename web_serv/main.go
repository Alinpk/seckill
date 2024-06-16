package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"goods_web/middleware"
	"goods_web/router"
)

func main() {

	//// 手动创建client
	//client := grpc.NewClient()

	r := gin.Default()

	// 解决跨域请求问题
	r.Use(middleware.CrosMiddleWare)
	// 加入日志中间件
	r.Use(middleware.Log())

	// 初始化router
	router.InitRouter(r)

	//r.GET("/test", func(c *gin.Context) {
	//	// 调用用户服务
	//	userSrvService := goods_user_srv.NewGoodsUserSrvService("go.micro.service.goods_user_srv", client)
	//	response1, _ := userSrvService.Call(context.Background(), &goods_user_srv.Request{Name: "goods_user_srv"})
	//	log.Info(response1.Msg)
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": response1.Msg,
	//	})
	//})

	service := web.NewService(
		web.Name("go.micro.service.goods_web"),
		web.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
		web.Version("latest"),
		web.Address(":9090"),
		web.Handler(r),
	)

	// 初始化
	service.Init()

	/*
		// New Service
		service := micro.NewService(
			micro.Name("go.micro.service.goods_web"),
			micro.Version("latest"),
		)

		// Initialise service
		service.Init()

		// Register Handler
		// goods_web.RegisterGoodsWebHandler(service.Server(), new(handler.Goods_web))
		goods_web.RegisterGoodsWebHandler(service.Server(), handler.NewGoodsWeb(service.Client()))

		// Register Struct as Subscriber
		//micro.RegisterSubscriber("go.micro.service.goods_web", service.Server(), new(subscriber.Goods_web))
	*/

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
