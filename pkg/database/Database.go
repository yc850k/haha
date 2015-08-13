package database

import (
	"gitlab.dian.fm/livecloud/config-server/pkg/dbstructs"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

type DB interface {
	Get(id int64) *DBStructs.Server
	//GetAll() []*DBStructs.Server
	Find(roomId int64) []*DBStructs.Server
	//Add(a *DBStructs.Server) (int, error)
	//Update(a *DBStructs.Server) error
	//Delete(id int64)
}

type serverDB struct {
	o  orm.Ormer
	qs orm.QuerySeter
}

func NewServerDB() serverDB {
	instance := orm.NewOrm()
	return serverDB{o: instance, qs: instance.QueryTable(new(DBStructs.Server))}
}

func init() {
	orm.RegisterModel(new(DBStructs.Server))
	orm.RegisterDataBase("default", "mysql", "root:hdoperation@tcp(yj-test-mysql-slb0:3306)/livecloud?charset=utf8", 30)
}

func (this *serverDB) Get(id int64) *DBStructs.Server {
	fmt.Printf("get mysql server: %v\n", id)
	server := DBStructs.Server{Id: id}
	err := this.o.Read(&server)
	fmt.Printf("ERR: %v\n", err)
	return &server
}

func (this *serverDB) Find(roomId int64) []*DBStructs.Server {
	var servers []*DBStructs.Server
	num, err := this.qs.Filter("room_id", roomId).All(&servers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return servers
}
