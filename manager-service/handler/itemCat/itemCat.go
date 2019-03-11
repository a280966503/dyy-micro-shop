package itemCat

import (
	"dyy-micro-shop/manager-web/proto/itemCat"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type ItemCat struct {
}

func (h *ItemCat) FindByParentId(ctx context.Context, in *itemCat.ItemCatModel, out *itemCat.ItemCatArr) error {

	var models []*itemCat.ItemCatModel
	o := orm.NewOrm()
	_, err := o.Raw("SELECT id,parent_id,name,type_id FROM tb_item_cat WHERE parent_id=?", in.ParentId).QueryRows(&models)

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out=itemCat.ItemCatArr{ItemCats:models}

	return nil
}

func (h *ItemCat) FindAll(ctx context.Context, in *itemCat.ItemCatModel, out *itemCat.ItemCatArr) error {

	var models []*itemCat.ItemCatModel
	o := orm.NewOrm()
	_, err := o.Raw("SELECT id,name FROM tb_item_cat").QueryRows(&models)

	if err != nil {
		fmt.Println(err)
		return err
	}

	*out=itemCat.ItemCatArr{ItemCats:models}
	return nil
}
