package mongodb

import "gopkg.in/mgo.v2"

const (
	URL="192.168.32.207:27017"
	C_ITEM="item"
)



func ConnectMongodb() (*mgo.Database,*mgo.Session)  {
	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("dyydb") //数据库名称
	return db,session
}