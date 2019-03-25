package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ShopLoginInfo struct {
	Username 	string
	Token 		string
	TokenTime	int64
}




type TbBrand struct {
	Id			int64			`json:"id"`
	Name 		string		`json:"name"`
	FirstChar	string		`json:"firstChar"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.3.4:3306)/pinyougoudb?charset=utf8")

	// register model
	orm.RegisterModel(
		new(TbBrand),
	)

	// create table
	orm.RunSyncdb("default", false, true)
}
