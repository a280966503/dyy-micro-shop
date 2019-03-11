package typeTemplate

import (
	"dyy-micro-shop/manager-web/proto/typeTemplate"
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
page=1&rows=10

--response
response.rows;//结果
response.total;//更新总记录数
{rows:[{}],total:10}
 */
func Search(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	page := utils.StringToInt32(r.Form.Get("page"))
	rows := utils.StringToInt32(r.Form.Get("rows"))

	service := micro.NewService()

	service.Init()

	templateService := typeTemplate.NewTypeTemplateService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respTemplate, err := templateService.Search(
		context.TODO(),
		&typeTemplate.ReqSearch{Page: page, Rows: rows},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respTemplate)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
{"customAttributeItems":[{"text":"1"},{"text":"1"}],"name":"1","brandIds":[{"id":2,"text":"华为"}],"specIds":[{"id":33,"text":"电视屏幕尺寸1"}]}

--response
{flag:true,message:""}
 */
func Add(w http.ResponseWriter,r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)


	var info map[string]interface{}


	json.Unmarshal(body,&info)

	//把json转string
	name := info["name"].(string)
	specIds := utils.JsonToString(info["specIds"])
	brandIds := utils.JsonToString(info["brandIds"])
	customAttributeItems := utils.JsonToString(info["customAttributeItems"])

	service := micro.NewService()

	service.Init()

	templateService := typeTemplate.NewTypeTemplateService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respTemplate, err := templateService.Add(
		context.TODO(),
		&typeTemplate.Rows{Name:name,SpecIds:specIds,BrandIds:brandIds,CustomAttributeItems:customAttributeItems},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respTemplate)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
ids=42,43

--response
{flag:true,message:""}
 */
func Delete(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	ids:= strings.Split(r.Form.Get("ids"),",")

	service := micro.NewService()

	service.Init()

	templateService := typeTemplate.NewTypeTemplateService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respTemplate, err := templateService.Delete(
		context.TODO(),
		&typeTemplate.ReqIds{Ids: ids},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respTemplate)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
?id=35
--response
Rows
 */
func FindOne(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	id := utils.StringToInt64(r.Form.Get("id"))

	service := micro.NewService()

	service.Init()

	templateService := typeTemplate.NewTypeTemplateService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respTemplate, err := templateService.FindOne(
		context.TODO(),
		&typeTemplate.Rows{Id: id},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respTemplate)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
{"id":1,"customAttributeItems":[{"text":"1"},{"text":"1"}],"name":"1","brandIds":[{"id":2,"text":"华为"}],"specIds":[{"id":33,"text":"电视屏幕尺寸1"}]}

--response
{flag:true,message:""}
 */
func Update(w http.ResponseWriter,r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)

	var info map[string]interface{}


	json.Unmarshal(body,&info)

	//把json转string
	id := int64(info["id"].(float64))
	name := info["name"].(string)
	specIds := utils.JsonToString(info["specIds"])
	brandIds := utils.JsonToString(info["brandIds"])
	customAttributeItems := utils.JsonToString(info["customAttributeItems"])

	service := micro.NewService()

	service.Init()

	templateService := typeTemplate.NewTypeTemplateService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respTemplate, err := templateService.Update(
		context.TODO(),
		&typeTemplate.Rows{Id:id,Name:name,SpecIds:specIds,BrandIds:brandIds,CustomAttributeItems:customAttributeItems},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respTemplate)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}