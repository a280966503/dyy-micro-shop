package seller

import (
	"dyy-micro-shop/manager-service/models"
	"dyy-micro-shop/manager-service/proto/seller"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type Seller struct {}

/**
--response
response.rows;//结果
response.total;//更新总记录数
 */
//seller_id,name,nick_name,password,email ,mobile ,telephone ,status ,address_detail ,linkman_name ,linkman_qq ,linkman_mobile ,linkman_email ,license_number ,tax_number ,org_number,address,logo_pic,brief,create_time,legal_person,legal_person_card_id,bank_user,bank_name
//调用微服务操作
func (c *Seller) Search(ctx context.Context, in *seller.ReqSeller, out *seller.RespSeller) error {

	page := in.Page
	rows := in.Rows

	status := in.Status
	start := (page-1)*rows

	var sellers []models.TbSeller

	o := orm.NewOrm()

	_, err := o.Raw(
		"select "+
			"seller_id,name,nick_name,password,email ,mobile ,telephone ,status ,address_detail ,linkman_name ,linkman_qq ,linkman_mobile ,linkman_email ,license_number ,tax_number ,org_number,address,logo_pic,brief,create_time,legal_person,legal_person_card_id,bank_user,bank_name "+
			"from tb_seller where status = ? limit ?,?", status, start, rows).QueryRows(&sellers)

	var count int64
	o.Raw("select count(seller_id) from tb_seller").QueryRow(&count)

	if err != nil {
		fmt.Println("查询失败:",err)
		return err
	}

	//回复数据
	*out = seller.RespSeller{Total:count}

	for _,value := range sellers{

		seller_rows := seller.Rows{
			SellerId:			value.SellerId          ,
			Name:				value.Name              ,
			NickName:			value.NickName          ,
			Password:			value.Password          ,
			Email:				value.Email             ,
			Mobile:				value.Mobile            ,
			Telephone:			value.Telephone         ,
			Status:				value.Status            ,
			AddressDetail:		value.AddressDetail     ,
			LinkmanName:		value.LinkmanName       ,
			LinkmanQq:			value.LinkmanQq         ,
			LinkmanMobile:		value.LinkmanMobile     ,
			LinkmanEmail:		value.LinkmanEmail      ,
			LicenseNumber:		value.LicenseNumber     ,
			TaxNumber:			value.TaxNumber         ,
			OrgNumber:			value.OrgNumber         ,
			Address:			value.Address          ,
			LogoPic:			value.LogoPic           ,
			Brief:				value.Brief             ,
			CreateTime:			""      ,
			LegalPerson:		value.LegalPerson       ,
			LegalPersonCardId:	value.LegalPersonCardId ,
			BankUser:			value.BankUser          ,
			BankName:			value.BankName          ,
		}

		out.Rows = append(out.Rows,&seller_rows)
	}

	return nil
}

func (c *Seller) FindOne(ctx context.Context, in *seller.ReqId, out *seller.Rows) error {

	var value models.TbSeller

	o := orm.NewOrm()

	o.Raw("select * from tb_seller where seller_id=?", in.Id).QueryRow(&value)

	//回复数据
	*out = seller.Rows{
		SellerId:			value.SellerId          ,
		Name:				value.Name              ,
		NickName:			value.NickName          ,
		Password:			value.Password          ,
		Email:				value.Email             ,
		Mobile:				value.Mobile            ,
		Telephone:			value.Telephone         ,
		Status:				value.Status            ,
		AddressDetail:		value.AddressDetail     ,
		LinkmanName:		value.LinkmanName       ,
		LinkmanQq:			value.LinkmanQq         ,
		LinkmanMobile:		value.LinkmanMobile     ,
		LinkmanEmail:		value.LinkmanEmail      ,
		LicenseNumber:		value.LicenseNumber     ,
		TaxNumber:			value.TaxNumber         ,
		OrgNumber:			value.OrgNumber         ,
		Address:			value.Address          ,
		LogoPic:			value.LogoPic           ,
		Brief:				value.Brief             ,
		CreateTime:			""      ,
		LegalPerson:		value.LegalPerson       ,
		LegalPersonCardId:	value.LegalPersonCardId ,
		BankUser:			value.BankUser          ,
		BankName:			value.BankName          ,
	}

	return nil

}

//WHERE LastName = 'Wilson'
func (c *Seller) UpdateStatus(ctx context.Context, in *seller.ReqIdAndStatus, out *seller.RespReturn) error {
	o := orm.NewOrm()

	fmt.Println(in)
	p, err := o.Raw("UPDATE tb_seller SET status = ? WHERE seller_id = ?").Prepare()

	if err != nil {
		*out=seller.RespReturn{Flag:false,Message:"更新失败"}
	}

	_, err = p.Exec(in.Status, in.SellerId)
	if err != nil {
		*out=seller.RespReturn{Flag:false,Message:"更新失败"}
	}

	*out = seller.RespReturn{Flag:true,Message:"更新成功"}

	return nil
}