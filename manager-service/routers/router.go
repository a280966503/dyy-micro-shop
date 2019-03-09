package routers

import (
	"dyy-micro-shop/manager-service/handler/brand"
	"dyy-micro-shop/manager-service/handler/seller"
	"dyy-micro-shop/manager-service/handler/specification"
	"github.com/micro/go-micro/server"
)

func Init(server server.Server) {

	/**********************seller**********************/
	server.Handle(server.NewHandler(&seller.Seller{}))
	/**********************brand**********************/
	server.Handle(server.NewHandler(&brand.Brand{}))
	/**********************specification**********************/
	server.Handle(server.NewHandler(&specification.Specification{}))
}
