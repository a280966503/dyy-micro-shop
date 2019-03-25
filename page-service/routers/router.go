package routers

import (
	"dyy-micro-shop/shop-service/handler/login"
	"github.com/micro/go-micro/server"
)

func Init(server server.Server) {

	/**********************login**********************/
	server.Handle(server.NewHandler(&login.Login{}))



}
