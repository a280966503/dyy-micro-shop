package main

import (
	"dyy-micro-shop/page-web/routers"
	"dyy-micro-shop/utils"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-web"
	"log"
)

func main() {

	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.32.204:2379",
		}
	})

	//设置端口创建服务
	service := web.NewService(
		web.Name(utils.SERVICE_PAGE_WEB),
		web.Registry(reg),
		web.Address(utils.ServicePort(utils.SERVICE_PAGE_WEB)),
	)

	//添加路由
	routers.Init(service)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

