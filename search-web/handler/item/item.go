package item

import (
	"dyy-micro-shop/common/proto/item"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
)

/**
--request
{"keywords":"dfds","category":"","brand":"","spec":{},"price":"","pageNo":1,"pageSize":40,"sort":"","sortField":""}
--response
{rows:[{}],totalPages:2}
 */
func Search(w http.ResponseWriter,r *http.Request)  {

	body, _ := ioutil.ReadAll(r.Body)

	var itemModel item.SearchParams
	json.Unmarshal(body,&itemModel)

	service := micro.NewService()
	service.Init()

	itemService := item.NewItemService(
		utils.SERVICE_CONTENT_SERVICE,
		service.Client(),
	)


	respSearch, err := itemService.Search(
		context.TODO(),
		&itemModel,
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respSearch)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}