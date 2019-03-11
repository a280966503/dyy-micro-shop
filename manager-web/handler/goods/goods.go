package goods

import (
	"dyy-micro-shop/manager-web/proto/goods"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

/**
--request
page=1&rows=10

--response
{total:100,rows:[{id:1,name:234,firstChar:N,firstName:2}]}
 */

func Search(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	page := utils.StringToInt32(r.Form.Get("page"))
	rows := utils.StringToInt32(r.Form.Get("rows"))

	service := micro.NewService()
	service.Init()

	goodsService := goods.NewGoodsService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respGoods, err := goodsService.Search(
		context.TODO(),
		&goods.ReqSearch{Page: page, Rows: rows},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respGoods)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
ids=149187842867912,149187842867915

--response
{flag:true,message:""}
 */
func Delete(w http.ResponseWriter,r *http.Request)  {

	r.ParseForm()
	ids := strings.Split(r.Form.Get("ids"), ",")

	service := micro.NewService()

	service.Init()

	goodsService := goods.NewGoodsService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respGoods, err := goodsService.Delete(
		context.TODO(),
		&goods.Ids{Ids: ids},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respGoods)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
ids=149187842867918,149187842867919&status=1
--response
{flag:true,message:""}
 */
func UpdateStatus(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	ids := strings.Split(r.Form.Get("ids"), ",")
	status := r.Form.Get("status")

	service := micro.NewService()

	service.Init()

	goodsService := goods.NewGoodsService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respGoods, err := goodsService.UpdateStatus(
		context.TODO(),
		&goods.Ids{Ids: ids,Status:status},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respGoods)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}