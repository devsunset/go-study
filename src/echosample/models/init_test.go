package models

import (
	"context"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
)

var ctx context.Context

func init() {
	runtime.GOMAXPROCS(1)
	xormEngine, err := xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	xormEngine.ShowSQL(true)
	xormEngine.Sync(new(Discount))
	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
}
