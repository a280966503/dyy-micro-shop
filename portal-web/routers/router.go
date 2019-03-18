package routers

import (
	"dyy-micro-shop/portal-web/handler/content"
	"github.com/micro/go-web"
	"net/http"
)

func Init(service web.Service) {
	//service.Handle("/",http.FileServer(http.Dir("views")))
	service.Handle("/",http.FileServer(http.Dir("views")))

	/*-----------------------content----------------------------*/
	service.Handle("/content/findByCategoryId.do",http.HandlerFunc(content.FindByCategoryId))


}
