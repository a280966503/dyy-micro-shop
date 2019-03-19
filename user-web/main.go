package main

import (
"dyy-micro-shop/shop-web/routers"
"dyy-micro-shop/utils"
"github.com/micro/go-web"
"log"
)

func main() {
	//设置端口创建服务
	service := web.NewService(
		web.Name(utils.SERVICE_USER_WEB),
		web.Address(utils.ServicePort(utils.SERVICE_USER_WEB)),
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

