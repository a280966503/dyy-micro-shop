package routers

import (
	"dyy-micro-shop/manager-web/handler/brand"
	"dyy-micro-shop/manager-web/handler/seller"
	"github.com/micro/go-web"
	"net/http"
)

func Init(service web.Service) {
	service.Handle("/",http.FileServer(http.Dir("views")))

	/*-----------------------seller----------------------------*/
	service.Handle("/seller/search.do",http.HandlerFunc(seller.Search))
	service.Handle("/seller/findOne.do",http.HandlerFunc(seller.FindOne))
	service.Handle("/seller/updateStatus.do",http.HandlerFunc(seller.UpdateStatus))

	/*-----------------------brand----------------------------*/
	service.Handle("/brand/search.do",http.HandlerFunc(brand.Search))
	service.Handle("/brand/save.do",http.HandlerFunc(brand.Save))
	service.Handle("/brand/findById.do",http.HandlerFunc(brand.FindById))
	service.Handle("/brand/delete.do",http.HandlerFunc(brand.Delete))
	service.Handle("/brand/update.do",http.HandlerFunc(brand.Update))



}
