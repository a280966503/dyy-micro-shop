package specification

import (
	"dyy-micro-shop/manager-web/proto/specification"
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

	rows := utils.StringToInt32(r.Form.Get("rows"))
	page := utils.StringToInt32(r.Form.Get("page"))

	service := micro.NewService()
	service.Init()

	specificationService := specification.NewSpecificationService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSpecification, err := specificationService.Search(
		context.TODO(),
		&specification.ReqSearch{Page: page, Rows: rows},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSpecification)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
{"specificationOptionList":[{"optionName":"2","orders":"2"},{"optionName":"2","orders":"2"}],"specification":{"specName":"2"}}

--response
{flag:true,message:""}
 */
func Add(w http.ResponseWriter,r *http.Request)  {

	body, _ := ioutil.ReadAll(r.Body)

	var info specification.ReqAdd
	json.Unmarshal(body,&info)

	service := micro.NewService()
	service.Init()

	specificationService := specification.NewSpecificationService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSpecification, err := specificationService.Add(
		context.TODO(),
		&info,
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSpecification)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
ids=26,27

--response
{flag:true}
 */
func Delete(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	ids := strings.Split(r.Form.Get("ids"), ",")

	service := micro.NewService()

	service.Init()

	specificationService := specification.NewSpecificationService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSpecification, err := specificationService.Delete(
		context.TODO(),
		&specification.ReqIds{Ids: ids},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSpecification)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
id=33

--response
{"specificationOptionList":[{"id":11,"specId":11,"optionName":"2","orders":"2"},{"id":11,"specId":11,"optionName":"2","orders":"2"}],"specification":{"id":11,"specName":"2"}}
 */
func FindOne(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	id := utils.StringToInt64(r.Form.Get("id"))

	service := micro.NewService()

	service.Init()

	specificationService := specification.NewSpecificationService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSpecification, err := specificationService.FindOne(
		context.TODO(),
		&specification.Rows{Id: id},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSpecification)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/**
--request
{"specification":{"id":26,"specName":"尺码"},"specificationOptionList":[{"id":91,"optionName":"165","specId":26,"orders":"1"},{"id":92,"optionName":"170","specId":26,"orders":"2"}]}
--response
{flag:true,message:""}
*/
func Update(w http.ResponseWriter,r *http.Request)  {

	var value specification.ReqAdd

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body,&value)

	service := micro.NewService()

	service.Init()

	specificationService := specification.NewSpecificationService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respSpecification, err := specificationService.Update(
		context.TODO(),
		&value,
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSpecification)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}