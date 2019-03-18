package content

import (
	"dyy-micro-shop/common/proto/content"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"net/http"
)

/**
--request
categoryId=1


[{}]
 */
func FindByCategoryId(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	categoryId := utils.StringToInt64(r.Form.Get("categoryId"))

	service := micro.NewService()

	service.Init()

	contentService := content.NewContentService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respContent, err := contentService.FindByCategoryId(
		context.TODO(),
		&content.ContentModel{CategoryId: categoryId},
	)

	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respContent.Rows)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}