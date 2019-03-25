package routers

import (
	"dyy-micro-shop/page-web/handler/item"
	"github.com/micro/go-web"
	"net/http"
)

func Init(service web.Service) {
	service.Handle("/", http.FileServer(http.Dir("views")))

	/*-----------------------login----------------------------*/
	service.Handle("/page.html", http.HandlerFunc(item.GetItem))

}