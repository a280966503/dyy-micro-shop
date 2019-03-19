package routers

import (
	"dyy-micro-shop/search-web/handler/item"
	"fmt"
	"github.com/micro/go-web"
	"net/http"
)

func Init(service web.Service) {
	//service.Handle("/",http.FileServer(http.Dir("views")))
	service.Handle("/",http.FileServer(http.Dir("views")))

	fmt.Println("===========================")

	/*-----------------------itemsearch----------------------------*/
	service.Handle("/itemsearch/search.do",http.HandlerFunc(item.Search))

}
