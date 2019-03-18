package login

import (
	"dyy-micro-shop/shop-web/proto/login"
	"dyy-micro-shop/utils"
	"encoding/json"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
--request
username=11111111&password=1&m1=2  post


 */
func CheckLogin(w http.ResponseWriter,r *http.Request)  {

	body, _ := ioutil.ReadAll(r.Body)
	//拆分字符串,并添加到map
	args := strings.Split(string(body[:]),"&")

	info := make(map[string]string)

	for _,value := range args{
		split := strings.Split(value, "=")
		key := split[0]
		info[key]=split[1]
	}

	service := micro.NewService()

	service.Init()

	loginService := login.NewLoginService(
		utils.SERVICE_SHOP_SERVICE,
		service.Client(),
	)

	respLogin, err := loginService.Login(
		context.TODO(),
		&login.Info{Username: info["username"], Password: info["password"]},
	)

	if err != nil {
		return
	}

	if respLogin.Flag {

	}else {

	}

	//转码为json
	resp, _ := json.Marshal(respLogin)

	//告诉前端数据类型
	w.Header().Set("Content-Type", "application/json")
	//会写数据
	w.Write(resp)

}