package main

import (
	"dyy-micro-shop/shop-service/routers"
	"dyy-micro-shop/utils"
	"fmt"
	"github.com/micro/go-micro"
	_ "dyy-micro-shop/shop-service/models"
)

func main() {
	//设置端口创建服务
	service := micro.NewService(
		micro.Name(utils.SERVICE_SHOP_SERVICE),
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

