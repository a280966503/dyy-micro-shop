package routers

import (
	"dyy-micro-shop/manager-service/handler/brand"
	"dyy-micro-shop/manager-service/handler/seller"
	"github.com/micro/go-micro/server"
)

func Init(server server.Server) {

	/**********************seller**********************/
	server.Handle(server.NewHandler(&seller.Seller{}))
	/**********************brand**********************/
	server.Handle(server.NewHandler(&brand.Brand{}))
}
