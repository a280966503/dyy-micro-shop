package login

import (
	"crypto/md5"
	"dyy-micro-shop/shop-service/models"
	"dyy-micro-shop/shop-web/proto/login"
	"dyy-micro-shop/utils"
	"dyy-micro-shop/utils/mongodb"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Login struct {
}


/**
--request
username=11111111&password=1&m1=2  post


 */
func (c *Login) Login(ctx context.Context, in *login.Info, out *login.Resp) error {

	fmt.Println("---------------")
	var username string
	o := orm.NewOrm()
	o.Raw("SELECT seller_id FROM tb_seller WHERE seller_id=? AND password=?", in.Username, in.Password).QueryRow(&username)

	fmt.Println("username:",username)
	if username!="" {
		//生成token
		crutime := string(time.Now().Unix())
		h := md5.New()
		h.Write([]byte(username+crutime)) // 需要加密的字符串为 123456
		cipherStr := h.Sum(nil)
		token := hex.EncodeToString(cipherStr)

		//token过期时间
		tokenTime := time.Now().Unix()+30*60

		//保存到mongodb中

		info := models.ShopLoginInfo{
			Username:username,
			Token:token,
			TokenTime:tokenTime,
		}
		db,session := mongodb.ConnectMongodb()
		defer session.Close()

		collection := db.C("shop_login_info")

		_, err := collection.RemoveAll(bson.M{"username": username})
		if err!=nil {
			fmt.Println(err)
		}

		err = collection.Insert(info)
		if err != nil {
			fmt.Println(err)
		}

		*out=login.Resp{Flag:true,Username:username,Token:token,Message:utils.RecodeText(utils.RECODE_OK)}
	}else {
		*out=login.Resp{Flag:false,Username:"",Token:"",Message:utils.RecodeText(utils.RECODE_LOGINERR)}
	}

	return nil
}