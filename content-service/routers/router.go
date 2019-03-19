package routers

import (
	"dyy-micro-shop/content-service/handler/item"
	"github.com/micro/go-micro/server"
)

func Init(server server.Server) {

	/**********************item**********************/
	server.Handle(server.NewHandler(&item.Item{}))

}
