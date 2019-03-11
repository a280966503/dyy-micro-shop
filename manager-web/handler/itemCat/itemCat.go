package itemCat

import (
	"dyy-micro-shop/manager-web/proto/itemCat"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"net/http"
)

/**
--request
?parentId=0
itemCatArr
 */
func FindByParentId(w http.ResponseWriter,r *http.Request)  {

	r.ParseForm()

	parentId := utils.StringToInt64(r.Form.Get("parentId"))

	service := micro.NewService()
	service.Init()

	catService := itemCat.NewItemCatService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respItemCat, err := catService.FindByParentId(
		context.TODO(),
		&itemCat.ItemCatModel{ParentId: parentId},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respItemCat.ItemCats)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}
/*
--request
[{"id":10,"name":""}]
*/
func FindAll(w http.ResponseWriter,r *http.Request)  {

	service := micro.NewService()

	service.Init()

	catService := itemCat.NewItemCatService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respItemCat, err := catService.FindAll(
		context.TODO(),
		&itemCat.ItemCatModel{},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respItemCat.ItemCats)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}