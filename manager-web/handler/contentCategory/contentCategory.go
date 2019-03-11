package contentCategory

import (
	"dyy-micro-shop/manager-web/proto/contentCategory"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
--request
?page=1&rows=10

--response
{total:100,rows:[{id:1,name:首页轮播广告}]}
 */
func Search(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	page := utils.StringToInt32(r.Form.Get("page"))
	rows := utils.StringToInt32(r.Form.Get("rows"))

	service := micro.NewService()
	service.Init()

	categoryService := contentCategory.NewContentCategoryService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respCategory, err := categoryService.Search(
		context.TODO(),
		&contentCategory.ReqContentCategory{Page: page, Rows: rows},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respCategory)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
{"name":"2"}

--response
{flag:true,message:""}
 */
func Add(w http.ResponseWriter,r *http.Request)  {

	body, _ := ioutil.ReadAll(r.Body)

	var model contentCategory.ContentCategoryModel
	json.Unmarshal(body,&model)

	service := micro.NewService()

	service.Init()

	categoryService := contentCategory.NewContentCategoryService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respCategory, err := categoryService.Add(
		context.TODO(),
		&contentCategory.ContentCategoryModel{Name: model.Name},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respCategory)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
ids=12,5

--response
{flag:true,message:""}

 */
func Delete(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	ids := strings.Split(r.Form.Get("ids"), ",")

	service := micro.NewService()

	service.Init()

	categoryService := contentCategory.NewContentCategoryService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respCategory, err := categoryService.Delete(
		context.TODO(),
		&contentCategory.Ids{Ids: ids},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respCategory)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
?id=1
 */
func FindOne(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	id := utils.StringToInt64(r.Form.Get("id"))

	service := micro.NewService()
	service.Init()

	categoryService := contentCategory.NewContentCategoryService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respCategory, err := categoryService.FindOne(
		context.TODO(),
		&contentCategory.ContentCategoryModel{Id: id},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respCategory)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request

 */
func FindAll(w http.ResponseWriter,r *http.Request)  {

	service := micro.NewService()
	service.Init()

	categoryService := contentCategory.NewContentCategoryService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respCategory, err := categoryService.FindAll(
		context.TODO(),
		&contentCategory.Req{},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respCategory.Rows)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}