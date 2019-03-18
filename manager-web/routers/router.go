package routers

import (
	"dyy-micro-shop/manager-web/handler/brand"
	"dyy-micro-shop/manager-web/handler/content"
	"dyy-micro-shop/manager-web/handler/contentCategory"
	"dyy-micro-shop/manager-web/handler/goods"
	"dyy-micro-shop/manager-web/handler/itemCat"
	login "dyy-micro-shop/manager-web/handler/login"
	"dyy-micro-shop/manager-web/handler/seller"
	"dyy-micro-shop/manager-web/handler/specification"
	"dyy-micro-shop/manager-web/handler/typeTemplate"
	"dyy-micro-shop/manager-web/handler/upload"
	"github.com/micro/go-web"
	"net/http"
)

func Init(service web.Service) {
	service.Handle("/",http.FileServer(http.Dir("views")))

	/*-----------------------seller----------------------------*/
	service.Handle("/seller/search.do",http.HandlerFunc(seller.Search))
	service.Handle("/seller/findOne.do",http.HandlerFunc(seller.FindOne))
	service.Handle("/seller/updateStatus.do",http.HandlerFunc(seller.UpdateStatus))

	/*-----------------------login----------------------------*/
	service.Handle("/login/search.do",http.HandlerFunc(brand.Search))
	service.Handle("/login/save.do",http.HandlerFunc(brand.Save))
	service.Handle("/login/findById.do",http.HandlerFunc(brand.FindById))
	service.Handle("/login/delete.do",http.HandlerFunc(brand.Delete))
	service.Handle("/login/update.do",http.HandlerFunc(brand.Update))
	service.Handle("/login/selectOptionList.do",http.HandlerFunc(brand.SelectOptionList))

	/*-----------------------specification----------------------------*/
	service.Handle("/specification/search.do",http.HandlerFunc(specification.Search))
	service.Handle("/specification/add.do",http.HandlerFunc(specification.Add))
	service.Handle("/specification/delete.do",http.HandlerFunc(specification.Delete))
	service.Handle("/specification/findOne.do",http.HandlerFunc(specification.FindOne))
	service.Handle("/specification/update.do",http.HandlerFunc(specification.Update))
	service.Handle("/specification/selectOptionList.do",http.HandlerFunc(specification.SelectOptionList))

	/*-----------------------typeTemplate----------------------------*/
	service.Handle("/typeTemplate/search.do",http.HandlerFunc(typeTemplate.Search))
	service.Handle("/typeTemplate/add.do",http.HandlerFunc(typeTemplate.Add))
	service.Handle("/typeTemplate/delete.do",http.HandlerFunc(typeTemplate.Delete))
	service.Handle("/typeTemplate/findOne.do",http.HandlerFunc(typeTemplate.FindOne))
	service.Handle("/typeTemplate/update.do",http.HandlerFunc(typeTemplate.Update))

	/*-----------------------itemCat----------------------------*/
	service.Handle("/itemCat/findByParentId.do",http.HandlerFunc(itemCat.FindByParentId))
	service.Handle("/itemCat/findAll.do",http.HandlerFunc(itemCat.FindAll))

	/*-----------------------goods----------------------------*/
	service.Handle("/goods/search.do",http.HandlerFunc(goods.Search))
	service.Handle("/goods/delete.do",http.HandlerFunc(goods.Delete))
	service.Handle("/goods/updateStatus.do",http.HandlerFunc(goods.UpdateStatus))

	/*-----------------------contentCategory----------------------------*/
	service.Handle("/contentCategory/search.do",http.HandlerFunc(contentCategory.Search))
	service.Handle("/contentCategory/add.do",http.HandlerFunc(contentCategory.Add))
	service.Handle("/contentCategory/delete.do",http.HandlerFunc(contentCategory.Delete))
	service.Handle("/contentCategory/findOne.do",http.HandlerFunc(contentCategory.FindOne))
	service.Handle("/contentCategory/findAll.do",http.HandlerFunc(contentCategory.FindAll))

	/*-----------------------content----------------------------*/
	service.Handle("/content/search.do",http.HandlerFunc(content.Search))
	service.Handle("/content/add.do",http.HandlerFunc(content.Add))
	service.Handle("/content/delete.do",http.HandlerFunc(content.Delete))
	service.Handle("/content/findOne.do",http.HandlerFunc(content.FindOne))
	service.Handle("/content/update.do",http.HandlerFunc(content.Update))

	/*-----------------------upload----------------------------*/
	service.Handle("/upload/uploadFile.do",http.HandlerFunc(upload.UploadFile))


	/*-----------------------manager----------------------------*/
	service.Handle("/manager/login",http.HandlerFunc(login.ManagerLogin))

}
