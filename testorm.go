package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"gitlab.dian.fm/livecloud/config-server/pkg/dbstructs"
)

func init() {
	fmt.Printf("init begin \n")
	// register model
	orm.RegisterModel(new(DBStructs.Server))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:hdoperation@tcp(yj-test-mysql-slb0:3306)/livecloud?charset=utf8", 30)
}

func main() {
	fmt.Printf("begin \n")
	o := orm.NewOrm()

	user := DBStructs.Server{Address: "tcp://127.0.0.1:8080"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Address = "udp://127.0.0.1:7001"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := DBStructs.Server{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
