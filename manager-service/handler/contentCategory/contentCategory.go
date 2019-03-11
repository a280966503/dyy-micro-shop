package contentCategory

import (
	"dyy-micro-shop/manager-web/proto/contentCategory"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type ContentCategory struct {

}

func (h *ContentCategory) Search(ctx context.Context, in *contentCategory.ReqContentCategory, out *contentCategory.RespContentCategory) error {

	var models []*contentCategory.ContentCategoryModel
	var count int64

	o := orm.NewOrm()

	_, err := o.Raw("SELECT id,name FROM tb_content_category LIMIT ?,?", (in.Page-1)*in.Rows, in.Rows).QueryRows(&models)
	o.Raw("SELECT COUNT(id) FROM tb_content_category").QueryRow(&count)

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out = contentCategory.RespContentCategory{Total:count,Rows:models}

	return nil
}

func (h *ContentCategory) Add(ctx context.Context, in *contentCategory.ContentCategoryModel, out *contentCategory.Resp) error {

	o := orm.NewOrm()

	_, err := o.Raw("INSERT INTO tb_content_category (name) values (?)", in.Name).Exec()

	if err!=nil {
		fmt.Println(err)
		return err
	}

	*out = contentCategory.Resp{Flag:true,Message:"添加成功"}
	return err
}


func (h *ContentCategory) Delete(ctx context.Context, in *contentCategory.Ids, out *contentCategory.Resp) error {

	o := orm.NewOrm()

	for _,id := range in.Ids{
		_, err := o.Raw("DELETE FROM tb_content_category WHERE id=?", id).Exec()
		if err!=nil {
			fmt.Println(err)
			return err
		}
	}
	*out=contentCategory.Resp{Flag:true,Message:"删除成功"}
	return nil
}

func (h *ContentCategory) FindOne(ctx context.Context, in *contentCategory.ContentCategoryModel, out *contentCategory.ContentCategoryModel) error {

	var model contentCategory.ContentCategoryModel
	o := orm.NewOrm()

	o.Raw("SELECT id,name FROM tb_content_category WHERE id=?",in.Id).QueryRow(&model)

	*out=model
	return nil
}

func (h *ContentCategory) FindAll(ctx context.Context, in *contentCategory.Req, out *contentCategory.RespContentCategory) error {

	var models []*contentCategory.ContentCategoryModel

	o := orm.NewOrm()

	_, err := o.Raw("SELECT id,name FROM tb_content_category").QueryRows(&models)

	if err!=nil {
		fmt.Println(err)
		return err
	}

	*out=contentCategory.RespContentCategory{Rows:models}
	return nil
}