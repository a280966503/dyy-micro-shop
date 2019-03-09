package seller

import (
	"context"
	seller "dyy-micro-shop/manager-web/proto/seller"
	utils "dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"io/ioutil"
	"net/http"
)


/**
--response
response.rows;//结果
response.total;//更新总记录数

--request
seller/search.do?page=1&rows=10
{status: "0"}
 */
func Search(w http.ResponseWriter,r *http.Request)  {
	//获取请求参数
	r.ParseForm()
	page := utils.StringToInt32(r.Form.Get("page"))
	rows := utils.StringToInt32(r.Form.Get("rows"))

	body, _ := ioutil.ReadAll(r.Body)

	var info map[string]string
	json.Unmarshal(body,&info)

	service := micro.NewService()

	//初始化
	service.Init()

	sellerService := seller.NewSellerService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client())


	//建立回复
	respSeller, err := sellerService.Search(
		context.TODO(),
		&seller.ReqSeller{Page: page, Rows: rows, Status: info["status"]})

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSeller)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/*
--request
id: alibaba

--response
response
*/
func FindOne(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	id := r.Form.Get("id")

	service := micro.NewService()
	service.Init()

	sellerService := seller.NewSellerService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSeller, err := sellerService.FindOne(
		context.TODO(),
		&seller.ReqId{Id: id},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSeller)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
sellerId=alibaba&status=1

--response
{flag:true,message:"成功"}
 */
func UpdateStatus(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	sellerId := r.Form.Get("sellerId")
	status := r.Form.Get("status")

	service := micro.NewService()
	service.Init()

	sellerService := seller.NewSellerService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSeller, err := sellerService.UpdateStatus(
		context.TODO(),
		&seller.ReqIdAndStatus{SellerId: sellerId, Status: status},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSeller)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}