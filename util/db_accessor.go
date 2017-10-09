package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func InitEngine() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/sports?charset=utf8")
	if err != nil {
		fmt.Print(err.Error)
		panic("Can not initialize database")
	}
}

func GetEngine() *xorm.Engine {
	return engine
}
