package routers

import (
	"dyy-micro-shop/shop-web/handler/goods"
	"dyy-micro-shop/shop-web/handler/itemCat"
	"dyy-micro-shop/shop-web/handler/login"
	"github.com/micro/go-web"
	"net/http"
)

func Init(service web.Service) {
	//service.Handle("/",http.FileServer(http.Dir("views")))
	service.Handle("/",http.FileServer(http.Dir("views")))

	/*-----------------------login----------------------------*/
	service.Handle("/shop/login",http.HandlerFunc(login.CheckLogin))

	/*-----------------------itemCat----------------------------*/
	//service.Handle("/itemCat/findOne.do",http.HandlerFunc(itemCat.FindOne))
	service.Handle("/itemCat/findAll.do",http.HandlerFunc(itemCat.FindAll))
	service.Handle("/itemCat/findByParentId.do",http.HandlerFunc(itemCat.FindByParentId))


	/*-----------------------goods----------------------------*/
	service.Handle("/goods/search.do",http.HandlerFunc(goods.Search))
}
