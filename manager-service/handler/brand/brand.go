package brand

import (
	"dyy-micro-shop/manager-service/models"
	"dyy-micro-shop/manager-web/proto/brand"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type Brand struct {

}
//id,name,first_char,first_name
func (h *Brand) Search(ctx context.Context, in *brand.ReqBrand, out *brand.RespBrand) error {



	page := in.Page
	rows := in.Rows

	firstChar := in.FirstChar
	name := in.Name

	var brands []models.TbBrand
	var count int64
	o := orm.NewOrm()
	if firstChar!="" && name!="" {
		o.Raw("SELECT id,name,first_char FROM tb_brand WHERE first_char=? AND name like ? LIMIT ?,?", firstChar, "%"+name+"%",(page-1)*rows,rows).
			QueryRows(&brands)
		o.Raw("SELECT count(id) FROM tb_brand WHERE first_char=? AND name like ?", firstChar, "%"+name+"%").
			QueryRow(&count)
	}else if firstChar != ""{
		o.Raw("SELECT id,name,first_char FROM tb_brand WHERE first_char=? LIMIT ?,?", firstChar,(page-1)*rows,rows).
			QueryRows(&brands)
		o.Raw("SELECT count(id) FROM tb_brand WHERE first_char=?", firstChar).
			QueryRow(&count)
	}else if name != ""{
		o.Raw("SELECT id,name,first_char FROM tb_brand WHERE name like ? LIMIT ?,?", "%"+name+"%",(page-1)*rows,rows).
			QueryRows(&brands)
		row := o.Raw("SELECT count(id) FROM tb_brand WHERE name like ?", "%"+name+"%").
			QueryRow(&count)
		fmt.Println("row",row)
	}else {
		o.Raw("SELECT id,name,first_char FROM tb_brand LIMIT ?,?",(page-1)*rows,rows).
			QueryRows(&brands)
		o.Raw("SELECT count(id) FROM tb_brand").
			QueryRow(&count)
	}
	fmt.Println(count)
	*out = brand.RespBrand{Total:count}

	for _,value := range brands{
		brand := brand.Rows{
			Id:        value.Id,
			Name:      value.Name,
			FirstChar: value.FirstChar,
		}
		out.Rows=append(out.Rows,&brand)
	}
	fmt.Println(out)
	return nil
}

//{flag:true,message:""}
func (c *Brand) Save(ctx context.Context, in *brand.Rows, out *brand.Resp) error {

	o := orm.NewOrm()

	_, err := o.Raw("INSERT INTO tb_brand(name,first_char) VALUES (?,?)", in.Name, in.FirstChar).Exec()

	if err != nil {
		fmt.Println(err)
	}

	*out = brand.Resp{Flag:true,Message:"添加成功"}

	return nil
}

//in--{id:xx,name:"",firstChar:""} -id 请求的参数 ,name和firstChar舍弃
func (h *Brand) FindById(ctx context.Context, in *brand.Rows, out *brand.Rows) error {

	var brandModel models.TbBrand
	o := orm.NewOrm()

	o.Raw("SELECT id,name,first_char FROM tb_brand WHERE id = ?",in.Id).QueryRow(&brandModel)

	*out = brand.Rows{
		Id:brandModel.Id,
		Name:brandModel.Name,
		FirstChar:brandModel.FirstChar,
	}
	return nil

}