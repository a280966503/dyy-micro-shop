package typeTemplate

import (
	"dyy-micro-shop/manager-service/proto/typeTemplate"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type TypeTemplate struct {
}

func (h *TypeTemplate) Search(ctx context.Context, in *typeTemplate.ReqSearch, out *typeTemplate.RespSearch) error {

	page := in.Page
	rows := in.Rows
	
	var models []*typeTemplate.Rows
	var count int64
	o := orm.NewOrm()

	_, err := o.Raw("SELECT id,name,spec_ids,brand_ids,custom_attribute_items,sspec_ids FROM tb_type_template LIMIT ?,?", (page-1)*rows, rows).QueryRows(&models)

	if err != nil {
		fmt.Println(err)
		return err
	}
	
	
	o.Raw("SELECT COUNT(id) FROM tb_type_template").QueryRow(&count)
	
	*out=typeTemplate.RespSearch{Total:count,Rows:models}

	return nil
}

func (h *TypeTemplate) Add(ctx context.Context, in *typeTemplate.Rows, out *typeTemplate.Resp) error {

	fmt.Println(in)
	name := in.Name
	specIds:= in.SpecIds
	brandIds:= in.BrandIds
	customAttributeItems := in.CustomAttributeItems
	sspecIds := in.SspecIds

	o := orm.NewOrm()

	_, err := o.Raw("INSERT INTO tb_type_template (name,spec_ids,brand_ids,custom_attribute_items,sspec_ids) VALUES (?,?,?,?,?)", name, specIds, brandIds, customAttributeItems, sspecIds).Exec()

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out = typeTemplate.Resp{Flag:true,Message:"添加成功"}

	return nil
}

func (h *TypeTemplate) Delete(ctx context.Context, in *typeTemplate.ReqIds, out *typeTemplate.Resp) error {

	o := orm.NewOrm()

	for _,id := range in.Ids{
		_, err := o.Raw("DELETE FROM tb_type_template WHERE id =?", id).Exec()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	*out =typeTemplate.Resp{Flag:true,Message:"删除成功"}

	return nil

}

func (h *TypeTemplate) FindOne(ctx context.Context, in *typeTemplate.Rows, out *typeTemplate.Rows) error {

	var model typeTemplate.Rows
	o := orm.NewOrm()

	o.Raw("SELECT id,name,spec_ids,brand_ids,custom_attribute_items FROM tb_type_template WHERE id=?", in.Id).QueryRow(&model)
	
	*out=model

	return nil
}

func (h *TypeTemplate) Update(ctx context.Context, in *typeTemplate.Rows, out *typeTemplate.Resp) error {

	fmt.Println(in)
	id := in.Id
	name := in.Name
	specIds:= in.SpecIds
	brandIds:= in.BrandIds
	customAttributeItems := in.CustomAttributeItems

	o := orm.NewOrm()

	_, err := o.Raw("UPDATE tb_type_template SET name=?,spec_ids=?,brand_ids=?,custom_attribute_items=? WHERE id=?", name, specIds, brandIds, customAttributeItems,id).Exec()

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out = typeTemplate.Resp{Flag:true,Message:"修改成功"}

	return nil
}

