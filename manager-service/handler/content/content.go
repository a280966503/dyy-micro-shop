package content

import (
	"dyy-micro-shop/manager-service/proto/content"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type Content struct {
}

func (h *Content) Search(ctx context.Context, in *content.ReqContent, out *content.RespContent) error {

	var models []*content.ContentModel
	var count int64

	o := orm.NewOrm()
	_, err := o.Raw("SELECT id,category_id,title,url,pic,status,sort_order FROM tb_content LIMIT ?,?", (in.Page-1)*in.Rows, in.Rows).QueryRows(&models)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	o.Raw("SELECT COUNT(id) FROM tb_content").QueryRow(&count)



	*out=content.RespContent{Total:count,Rows:models}
	return nil
}

func (h *Content) Add(ctx context.Context, in *content.ContentModel, out *content.Resp) error {

	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO tb_content (category_id,title,url,pic,status,sort_order) VALUES (?,?,?,?,?,?)", in.CategoryId, in.Title, in.Url, in.Pic, in.Status, in.SortOrder).Exec()

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out=content.Resp{Flag:true,Message:"添加成功"}
	return nil

}


func (h *Content) Delete(ctx context.Context, in *content.Ids, out *content.Resp) error {

	o := orm.NewOrm()

	for _,id := range in.Ids{
		_, err := o.Raw("DELETE FROM tb_content WHERE id=?", id).Exec()

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	*out=content.Resp{Flag:true,Message:"删除成功"}

	return nil
}

func (h *Content) FindOne(ctx context.Context, in *content.ContentModel, out *content.ContentModel) error {

	var model content.ContentModel

	o := orm.NewOrm()

	o.Raw("SELECT id,category_id,title,url,pic,status,sort_order FROM tb_content WHERE id=?",in.Id).QueryRow(&model)

	*out=model

	return nil
}

func (h *Content) Update(ctx context.Context, in *content.ContentModel, out *content.Resp) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE tb_content SET category_id=?,title=?,url=?,pic=?,status=?,sort_order=? WHERE id=?", in.CategoryId, in.Title, in.Url, in.Pic, in.Status, in.SortOrder, in.Id).Exec()

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out = content.Resp{Flag:true,Message:"修改成功"}
	return nil
}