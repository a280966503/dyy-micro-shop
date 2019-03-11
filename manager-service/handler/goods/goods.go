package goods

import (
	"dyy-micro-shop/manager-web/proto/goods"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type Goods struct {
}

const sqlName  = "id,seller_id,goods_name,default_item_id,audit_status,is_marketable,brand_id,caption,category1_id,category2_id,category3_id,small_pic,price,type_template_id,is_enable_spec,is_delete"

func (h *Goods) Search(ctx context.Context, in *goods.ReqSearch, out *goods.RespSearch) error {

	rows := in.Rows

	page := in.Page

	var models []*goods.GoodsModel
	var count int64
	o := orm.NewOrm()

	_, err := o.Raw("SELECT "+sqlName+" FROM tb_goods LIMIT ?,?", (page-1)*rows, rows).QueryRows(&models)
	if err != nil {
		fmt.Println(err)
		return err
	}
	o.Raw("SELECT COUNT(id) FROM tb_goods").QueryRow(&count)


	*out=goods.RespSearch{Total:count,Rows:models}
	return nil


}

func (h *Goods) Delete(ctx context.Context, in *goods.Ids, out *goods.Resp) error {

	o := orm.NewOrm()
	for _,id := range in.Ids{
		_, err := o.Raw("DELETE FROM tb_goods WHERE id=?", id).Exec()
		if err != nil {
			fmt.Println(err)

			return err
		}
	}
	*out = goods.Resp{Flag:true,Message:"删除成功"}
	return nil


}

func (h *Goods) UpdateStatus(ctx context.Context, in *goods.Ids, out *goods.Resp) error {
	o := orm.NewOrm()
	for _,id := range in.Ids{
		_, err := o.Raw("UPDATE tb_goods SET audit_status=? WHERE id=?",in.Status, id).Exec()
		if err != nil {
			fmt.Println(err)

			return err
		}
	}
	*out = goods.Resp{Flag:true,Message:"审核成功"}
	return nil
}
