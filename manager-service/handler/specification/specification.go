package specification

import (
	"dyy-micro-shop/manager-service/models"
	"dyy-micro-shop/manager-service/proto/specification"
	"dyy-micro-shop/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type Specification struct {
}

func (h *Specification) Search(ctx context.Context, in *specification.ReqSearch, out *specification.RespSearch) error {

	var specificationModels []models.TbSpecification
	var count int64
	o := orm.NewOrm()

	_, err := o.Raw("SELECT id,spec_name FROM tb_specification LIMIT ?,?",(in.Page-1)*in.Rows,in.Rows).QueryRows(&specificationModels)
	o.Raw("SELECT count(id) FROM tb_specification ").QueryRow(&count)


	if err != nil {
		fmt.Println(err)
		return err
	}

	*out=specification.RespSearch{Total:count}

	for _,value := range specificationModels{

		rows := specification.Rows{Id: value.Id, SpecName: value.SpecName}
		out.Rows = append(out.Rows,&rows)
	}

	return nil
}

/**
--request
{"specificationOptionList":[{"optionName":"2","orders":"2"},{"optionName":"2","orders":"2"}],"specification":{"specName":"2"}}

--response
{flag:true,message:""}
 */
func (h *Specification) Add(ctx context.Context, in *specification.ReqAdd, out *specification.Resp) error {

	specName := in.Specification.SpecName



	o := orm.NewOrm()

	o.Begin()
	result, err := o.Raw("INSERT INTO tb_specification (spec_name) VALUES (?)", specName).Exec()

	specId, _ := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	for _,specificationOption := range in.SpecificationOptionList {
		optionName:=specificationOption.OptionName
		orders:=utils.StringToInt64(specificationOption.Orders)

		_, err = o.Raw("INSERT INTO tb_specification_option (option_name,spec_id,orders) VALUES (?,?,?)",optionName,specId,orders).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	*out=specification.Resp{Flag:true,Message:"添加成功"}

	o.Commit()

	return nil

}


func (h *Specification) Delete(ctx context.Context, in *specification.ReqIds, out *specification.Resp) error {

	o := orm.NewOrm()

	o.Begin()

	for _,id := range in.Ids{
		_, err := o.Raw("DELETE FROM tb_specification WHERE id=?", id).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		_, err = o.Raw("DELETE FROM tb_specification_option WHERE spec_id=?", id).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	*out=specification.Resp{Flag:true}
	o.Commit()
	return nil
}

func (h *Specification) FindOne(ctx context.Context, in *specification.Rows, out *specification.ReqAdd) error {

	var specificationModel specification.Rows
	o := orm.NewOrm()

	o.Raw("SELECT id,spec_name FROM tb_specification WHERE id=?",in.Id).QueryRow(&specificationModel)

	var specificationOptionModels []*specification.ReqAdd_SpecificationOption
	_, err := o.Raw("SELECT id,option_name,spec_id,orders FROM tb_specification_option WHERE spec_id=? ", in.Id).QueryRows(&specificationOptionModels)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(specificationModel)

	*out=specification.ReqAdd{Specification:&specificationModel,SpecificationOptionList:specificationOptionModels}
	return nil
}

func (h *Specification) Update(ctx context.Context, in *specification.ReqAdd, out *specification.Resp) error {

	specificationModel:=in.Specification
	o := orm.NewOrm()

	o.Begin()
	_, err := o.Raw("UPDATE tb_specification SET spec_name=? WHERE id=? ", specificationModel.SpecName, specificationModel.Id).Exec()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	for _,specificationOption := range in.SpecificationOptionList{
		_, err := o.Raw("UPDATE tb_specification_option SET option_name=?,orders=?", specificationOption.OptionName, specificationOption.Orders).Exec()
		if err != nil {
			o.Rollback()
			fmt.Println(err)
			return err
		}

	}
	*out=specification.Resp{Flag:true,Message:"修改成功"}
	o.Commit()
	return nil

}
