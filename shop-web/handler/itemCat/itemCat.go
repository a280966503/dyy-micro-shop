package itemCat

import (
	"dyy-micro-shop/shop-web/proto/itemCat"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"net/http"
)

/*
--request
id=3

--response

*/
func FindOne(w http.ResponseWriter,r *http.Request)  {

	//r.ParseForm()
	//id := utils.StringToInt64(r.Form.Get("id"))


}

/**
--response
[{id:1,parentId:1,name:"eeee",typeId:1}]
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
	resp, _ := json.Marshal(respItemCat)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
parentId=undefined

--response
[{"id":1,"name":"图书、音像、电子书刊","typeId":35},{"id":74,"name":"家用电器","typeId":35},{"id":161,"name":"电脑、办公","typeId":35},{"id":249,"name":"个护化妆","typeId":35},{"id":290,"name":"钟表","typeId":35},{"id":296,"name":"母婴","typeId":35},{"id":378,"name":"食品饮料、保健食品","typeId":35},{"id":438,"name":"汽车用品","typeId":35},{"id":495,"name":"玩具乐器","typeId":35},{"id":558,"name":"手机","typeId":35},{"id":580,"name":"数码","typeId":35},{"id":633,"name":"家居家装","typeId":35},{"id":699,"name":"厨具","typeId":35},{"id":749,"name":"服饰内衣","typeId":35},{"id":865,"name":"鞋靴","typeId":35},{"id":903,"name":"礼品箱包","typeId":35},{"id":963,"name":"珠宝","typeId":35},{"id":1031,"name":"运动健康","typeId":35},{"id":1147,"name":"彩票、旅行、充值、票务","typeId":35},{"id":1186,"name":"小白大数据","typeId":35}]
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