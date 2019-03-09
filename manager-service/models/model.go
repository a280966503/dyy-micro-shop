package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//seller_id,name,nick_name,password,email ,mobile ,telephone ,status ,address_detail ,linkman_name ,linkman_qq ,linkman_mobile ,linkman_email ,license_number ,tax_number ,org_number,address,logo_pic,brief,create_time,legal_person,legal_person_card_id,bank_user,bank_name
type TbSeller struct {
	SellerId 				string 		`orm:"column(seller_id);pk" json:"sellerId"`
	Name 					string		`json:"name"`
	NickName 				string		`json:"nickName"`
	Password 				string		`json:"password"`
	Email 					string		`json:"email"`
	Mobile 					string		`json:"mobile"`
	Telephone 				string		`json:"telephone"`
	Status 					string		`json:"status"`
	AddressDetail 			string		`json:"addressDetail"`
	LinkmanName 			string		`json:"linkmanName"`
	LinkmanQq 				string		`json:"linkmanQq"`
	LinkmanMobile			string		`json:"linkmanMobile"`
	LinkmanEmail			string		`json:"linkmanEmail"`
	LicenseNumber			string		`json:"licenseNumber"`
	TaxNumber				string		`json:"taxNumber"`
	OrgNumber				string		`json:"orgNumber"`
	Address					int64		`json:"address"`
	LogoPic					string		`json:"logoPic"`
	Brief					string		`json:"brief"`
	CreateTime				time.Time	`json:"createTime"`
	LegalPerson				string		`json:"legalPerson"`
	LegalPersonCardId		string		`json:"legalPersonCardId"`
	BankUser				string		`json:"bankUser"`
	BankName				string		`json:"bankName"`
}
//{"specificationOptionList":[{"optionName":"16G","orders":"1"},{"optionName":"8G","orders":"2"}],"specification":{"specName":"内存"}}

type Spec struct {
	Specification 			*TbSpecification 			`json:"specification"`
	SpecificationOptionList []*TbSpecificationOption 	`json:"specificationOptionList"`
}

type TbSpecification struct {
	Id			int64		`json:"id"`
	SpecName 	string `json:"specName"`
}

type TbSpecificationOption struct {
	Id 			int		`json:"id"`
	SpecId 		int		`json:"specId"`
	OptionName 	string `json:"optionName"`
	Orders     	string `json:"orders"`
}

/*********************brand************************/
type TbBrand struct {
	Id			int64			`json:"id"`
	Name 		string		`json:"name"`
	FirstChar	string		`json:"firstChar"`
}

type TbContent struct {
	Id			int		`json:"id"`
	CategoryId 	int		`json:"categoryId"`
	Title		string	`json:"title"`
	Url 		string	`json:"url"`
	Pic 		string	`json:"pic"`
	Status		string	`json:"status"`
	SortOrder	string 	`json:"sortOrder"`

}

type TbContentCategory struct {
	Id			int 	`json:"id"`
	Name 		string	`json:"name"`
}

type GoodsAndDesc struct {
	Goods 		*TbGoods		`json:"goods"`
	GoodsDesc 	*TbGoodsDesc	`json:"goodsDesc"`
}


type TbGoods struct {
	Id 					uint64	`json:"id"`
	SellerId			string	`json:"sellerId"`
	GoodsName			string	`json:"goodsName"`
	DefaultItemId		int64		`json:"defaultItemId"`
	AuditStatus			string	`json:"auditStatus"`
	IsMarketable		string	`json:"isMarketable"`
	BrandId				int64		`json:"brandId"`
	Caption				string	`json:"caption"`
	Category1Id			int64		`json:"category1Id"`
	Category2Id			int64		`json:"category2Id"`
	Category3Id			int64		`json:"category3Id"`
	SmallPic			string	`json:"smallPic"`
	Price				float64	`json:"price"`
	TypeTemplateId		int64		`json:"typeTemplateId"`
	IsEnableSpec		string	`json:"isEnableSpec"`
	IsDelete			string	`json:"isDelete"`
}

type TbGoodsDesc struct {
	GoodsId					int64	`orm:"column(goods_id);pk" json:"goodsId"`
	Introduction			string	`json:"introduction"`
	SpecificationItems		string	`json:"specificationItems"`
	CustomAttributeItems	string	`json:"customAttributeItems"`
	ItemImages				string	`json:"itemImages"`
	PackageList				string	`json:"packageList"`
	SaleService				string	`json:"saleService"`
}



type TbItemCat struct {
	Id			int64	`json:"id"`
	ParentId	int64	`json:"parentId"`
	Name		string	`json:"name"`
	TypeId		int64	`json:"typeId"`
}

type TbTypeTemplate struct {
	Id						int64	`json:"id"`
	Name					string	`json:"name"`
	SpecIds					string	`json:"specIds"`
	BrandIds				string	`json:"brandIds"`
	CustomAttributeItems	string	`json:"customAttributeItems"`
}


func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.3.4:3306)/pinyougoudb?charset=utf8")

	// register model
	orm.RegisterModel(
		new(TbBrand),
		new(TbSeller),
		new(TbSpecification),
	)

	// create table
	orm.RunSyncdb("default", false, true)
}
