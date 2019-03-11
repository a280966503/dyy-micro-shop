package brand

import (
	"dyy-micro-shop/manager-web/proto/brand"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strings"
)


/*
--request
page=1&rows=10
{"name":"fghfg","firstChar":"fgdfg"} 状态可为1.{} 2.{"name":"fgsdf"} 3.{"firstChar":"wew"} 4.{"name":"fghfg","firstChar":"fgdfg"}这几种

--response
{total:100,rows:[{id:1,name:234,firstChar:N,firstName:2}]}
*/
func Search(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	page := utils.StringToInt32(r.Form.Get("page"))
	rows := utils.StringToInt32(r.Form.Get("rows"))

	body, _ := ioutil.ReadAll(r.Body)

	var info map[string]string

	json.Unmarshal(body,&info)

	service := micro.NewService()

	service.Init()

	brandService := brand.NewBrandService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respBrand, err := brandService.Search(
		context.TODO(),
		&brand.ReqBrand{Page: page, Rows: rows, Name: info["name"], FirstChar: info["firstChar"]},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respBrand)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/*
--request
{"name":"1","firstChar":"1"}

--response
{flag:true,message:""}
*/
func Save(w http.ResponseWriter,r *http.Request)  {

	body, _ := ioutil.ReadAll(r.Body)

	var info map[string]string

	json.Unmarshal(body,&info)

	service:= micro.NewService()

	service.Init()

	brandService := brand.NewBrandService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respBrand, err := brandService.Save(
		context.TODO(),
		&brand.Rows{Name: info["name"], FirstChar: info["firstChar"]},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respBrand)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}

/*
--request
id=1

--response
{id:xx,name:yy,firstChar:zz}

*/

func FindById(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	id := utils.StringToInt64(r.Form.Get("id"))

	service := micro.NewService()
	service.Init()

	brandService := brand.NewBrandService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respBrand, err := brandService.FindById(
		context.TODO(),
		&brand.Rows{Id: id},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respBrand)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
ids=1,2,3

--response
{flag:true,message:""}
 */
func Delete(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	ids := strings.Split(r.Form.Get("ids"),",")

	service := micro.NewService()

	service.Init()

	brandService := brand.NewBrandService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respBrand, err := brandService.Delete(
		context.TODO(),
		&brand.ReqIds{Ids: ids},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respBrand)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--request
{"id":21,"name":"康佳","firstChar":"KD"}

--response
{flag:true,message:""}
 */
func Update(w http.ResponseWriter,r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)

	var info brand.Rows

	json.Unmarshal(body,&info)

	service := micro.NewService()

	service.Init()

	brandService := brand.NewBrandService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respBrand, err := brandService.Update(
		context.TODO(),
		&info,
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respBrand)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}

/**
--response
[{"id":1,"text":"联想"},{"id":3,"text":"三星"}]
 */
func SelectOptionList(w http.ResponseWriter,r *http.Request)  {

	service := micro.NewService()

	service.Init()

	brandService := brand.NewBrandService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respBrand, err := brandService.SelectOptionList(
		context.TODO(),
		&brand.Req{},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respBrand.OptionList)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}