package routers

import (
	"dyy-micro-shop/manager-web/handler/brand"
	"dyy-micro-shop/manager-web/handler/seller"
	"dyy-micro-shop/manager-web/handler/specification"
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

	/*-----------------------specification----------------------------*/
	service.Handle("/specification/search.do",http.HandlerFunc(specification.Search))
	service.Handle("/specification/add.do",http.HandlerFunc(specification.Add))
	service.Handle("/specification/delete.do",http.HandlerFunc(specification.Delete))
	service.Handle("/specification/findOne.do",http.HandlerFunc(specification.FindOne))
	service.Handle("/specification/update.do",http.HandlerFunc(specification.Update))
}
