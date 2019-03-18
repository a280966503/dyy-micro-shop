package upload

import (
	"dyy-micro-shop/manager-web/proto/upload"
	"dyy-micro-shop/utils"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
)

/**
--request
[]byte

--response
{flag:true,message:""}
 */
func UploadFile(w http.ResponseWriter,r *http.Request)   {




	contentType := r.Header.Get("Content-Type")
	mediatype, params, err := mime.ParseMediaType(contentType)
	if mediatype != "multipart/form-data" {
		fmt.Println(err)
		return
	}
	boundary, _ := params["boundary"]
	reader := multipart.NewReader(r.Body, boundary)
	part, _ := reader.NextPart()


	bytes,_ := ioutil.ReadAll(part)
	service := micro.NewService()

	service.Init()

	uploadService := upload.NewUploadService(
		utils.SERVICE_MANAGER_SERVICE,
		service.Client(),
	)

	respUpload, err := uploadService.UploadFile(
		context.TODO(),
		&upload.ReqUpload{Img: bytes},
	)
	if err != nil {
		fmt.Println("=========",err)
		return
	}

	//转码为json
	resp, _ := json.Marshal(respUpload)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)
}