package main

import (
	"dyy-micro-shop/content-service/routers"
	"dyy-micro-shop/utils"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.32.204:2379",
		}
	})
	//设置端口创建服务
	service := micro.NewService(
		micro.Name(utils.SERVICE_CONTENT_SERVICE),
		micro.Registry(reg),
		micro.Version("latest"),
	)

	//初始化
	service.Init()
	//创建服务器句柄
	server := service.Server()

	routers.Init(server)

	err := service.Run()

	if err != nil {
		fmt.Println(err)
	}
}

