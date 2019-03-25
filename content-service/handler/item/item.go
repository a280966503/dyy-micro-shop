package item

import (
	"dyy-micro-shop/common/proto/item"
	"dyy-micro-shop/utils/mongodb"
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
}

//从mongodb中获取商品信息
//{"keywords":"dfds","category":"","brand":"","spec":{},"price":"","pageNo":1,"pageSize":40,"sort":"","sortField":""}

func (h *Item) Search(ctx context.Context, in *item.SearchParams, out *item.RespSearch) error {


	database, session := mongodb.ConnectMongodb()

	defer session.Close()

	var models []*item.Model

	collection := database.C(mongodb.C_ITEM)

	//查询分页商品
	//bson.M{"title": "new2 - 阿尔卡特 (OT-927) 炭黑 联通3G手机 双卡双待"}
	collection.Find(bson.M{"title": bson.M{"$regex":in.Keywords, "$options": "$i"}}).
		Limit(int(in.PageSize)).Skip(int(in.PageSize * (in.PageNo - 1))).All(&models)
	count, err := collection.Find(bson.M{"title": bson.M{"$regex": in.Keywords, "$options": "$i"}}).Count()

	if err != nil {
		fmt.Println(err)
		return err
	}


//item.RespSearch{Rows:models,TotalPages:int64(count)/in.PageSize+1,Total:int64(count)}
	*out = item.RespSearch{Rows:models,TotalPages:int64(count)/in.PageSize+1}

	return nil
}