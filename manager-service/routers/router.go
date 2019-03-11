package routers

import (
	"dyy-micro-shop/manager-service/handler/brand"
	"dyy-micro-shop/manager-service/handler/content"
	"dyy-micro-shop/manager-service/handler/contentCategory"
	"dyy-micro-shop/manager-service/handler/goods"
	"dyy-micro-shop/manager-service/handler/itemCat"
	"dyy-micro-shop/manager-service/handler/seller"
	"dyy-micro-shop/manager-service/handler/specification"
	"dyy-micro-shop/manager-service/handler/typeTemplate"
	"github.com/micro/go-micro/server"
)

func Init(server server.Server) {

	/**********************brand**********************/
	server.Handle(server.NewHandler(&brand.Brand{}))

	/**********************content**********************/
	server.Handle(server.NewHandler(&content.Content{}))

	/**********************contentCategory**********************/
	server.Handle(server.NewHandler(&contentCategory.ContentCategory{}))

	/**********************goods**********************/
	server.Handle(server.NewHandler(&goods.Goods{}))

	/**********************itemCat**********************/
	server.Handle(server.NewHandler(&itemCat.ItemCat{}))

	/**********************seller**********************/
	server.Handle(server.NewHandler(&seller.Seller{}))

	/**********************specification**********************/
	server.Handle(server.NewHandler(&specification.Specification{}))

	/**********************typeTemplate**********************/
	server.Handle(server.NewHandler(&typeTemplate.TypeTemplate{}))


}
