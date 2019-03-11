package content

import (
	"dyy-micro-shop/manager-web/proto/content"
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
{total:100,rows:[]}
 */
func Search(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	page := utils.StringToInt32(r.Form.Get("page"))
	rows := utils.StringToInt32(r.Form.Get("rows"))

	service := micro.NewService()
	service.Init()

	contentService := content.NewContentService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respContent, err := contentService.Search(
		context.TODO(),
		&content.ReqContent{Page: page, Rows: rows},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respContent)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
{"categoryId":1,"title":"1","url":"1","sortOrder":"1","status":"1"}

--response
{flag:true,message:""}
 */
func Add(w http.ResponseWriter,r *http.Request)  {
	var model content.ContentModel

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&model)

	service := micro.NewService()
	service.Init()

	contentService := content.NewContentService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respContent, err := contentService.Add(
		context.TODO(),
		&model,
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respContent)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
?ids=13

--response
{flag:true,message:""}
 */
func Delete(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	ids := strings.Split(r.Form.Get("ids"), ",")

	service := micro.NewService()

	service.Init()

	contentService := content.NewContentService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respContent, err := contentService.Delete(
		context.TODO(),
		&content.Ids{Ids: ids},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respContent)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
id=5

--response
{flag:true,message:""}

 */
func FindOne(w http.ResponseWriter,r *http.Request)  {

	r.ParseForm()
	id := utils.StringToInt64(r.Form.Get("id"))

	service := micro.NewService()
	service.Init()

	contentService := content.NewContentService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respContent, err := contentService.FindOne(
		context.TODO(),
		&content.ContentModel{Id: id},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respContent)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
{"id":5,"categoryId":1,"title":"测试广告222","url":"http://www.baidu.com","pic":"http://192.168.25.133/group1/M00/00/00/wKgZhVnIhVqAAa0jAActhhatATY291.png","status":"1","sortOrder":3}
--response
{flag:true,message:""}
 */

func Update(w http.ResponseWriter,r *http.Request)  {
	var model content.ContentModel

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&model)

	service := micro.NewService()

	service.Init()

	contentService := content.NewContentService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respContent, err := contentService.Update(
		context.TODO(),
		&model,
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respContent)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}