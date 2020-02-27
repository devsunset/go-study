package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wenj91/gobatis"
)

type User struct {
	Id    gobatis.NullInt64  `field:"id"`
	Name  gobatis.NullString `field:"name"`
	Email gobatis.NullString `field:"email"`
	CrtTm gobatis.NullTime   `field:"crtTm"`
}

func main() {
	example()
	example1()
	example2()
	example3()
}

func example() {
	db, _ := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	dbs := make(map[string]*gobatis.GoBatisDB)
	dbs["ds1"] = gobatis.NewGoBatisDB(gobatis.DBTypeMySQL, db)

	option := gobatis.NewDBOption().
		DB(dbs).
		ShowSQL(true).
		Mappers([]string{"mapper/userMapper.xml"})

	gobatis.Init(option)

	gb := gobatis.Get("ds1")

	mapRes := make(map[string]interface{})
	err := gb.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
	fmt.Println("userMapper.findMapById-->", mapRes, err)

	mapRes2 := make(map[string]interface{})
	err = gb.SelectContext(context.TODO(), "userMapper.findMapById", map[string]interface{}{"id": 4})(mapRes2)
	fmt.Println("userMapper.findMapById-->", mapRes2, err)

	fmt.Println("############################ --- example")
}

func example1() {

	gobatis.Init(gobatis.NewFileOption("db.yml"))

	// datasource1
	gb := gobatis.Get("ds1")

	mapRes := make(map[string]interface{})
	err := gb.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
	fmt.Println("userMapper.findMapById-->", mapRes, err)

	param := User{Id: gobatis.NullInt64{Int64: 1, Valid: true}}
	var structRes *User
	err = gb.Select("userMapper.findStructByStruct", param)(&structRes)
	fmt.Println("userMapper.findStructByStruct-->", structRes, err)

	structsRes := make([]*User, 0)
	err = gb.Select("userMapper.queryStructs", map[string]interface{}{})(&structsRes)
	fmt.Println("userMapper.queryStructs-->", structsRes, err)

	param = User{
		Id:   gobatis.NullInt64{Int64: 1, Valid: true},
		Name: gobatis.NullString{String: "wenj1993", Valid: true},
	}
	affected, err := gb.Update("userMapper.updateByCond", param)
	fmt.Println("updateByCond:", affected, err)

	param = User{Name: gobatis.NullString{String: "wenj1993", Valid: true}}
	res := make([]*User, 0)
	err = gb.Select("userMapper.queryStructsByCond", param)(&res)
	fmt.Println("queryStructsByCond", res, err)

	res = make([]*User, 0)
	err = gb.Select("userMapper.queryStructsByCond2", param)(&res)
	fmt.Println("queryStructsByCond2", res, err)

	res = make([]*User, 0)
	err = gb.Select("userMapper.queryStructsByOrder", map[string]interface{}{
		"id": "id",
	})(&res)
	fmt.Println("queryStructsByCond", res, err)

	tx, _ := gb.Begin()
	defer tx.Rollback()
	tx.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
	fmt.Println("tx userMapper.findMapById-->", mapRes, err)
	tx.Commit()

	fmt.Println("############################ --- example1")
}

func example2() {
	ds1 := gobatis.NewDataSourceBuilder().
		DataSource("ds1").
		DriverName("mysql").
		DataSourceName("root:admin@tcp(127.0.0.1:3306)/test?charset=utf8").
		MaxLifeTime(120).
		MaxOpenConns(10).
		MaxIdleConns(5).
		Build()

	option := gobatis.NewDSOption().
		DS([]*gobatis.DataSource{ds1}).
		Mappers([]string{"mapper/userMapper.xml"}).
		ShowSQL(true)

	gobatis.Init(option)

	gb := gobatis.Get("ds1")

	mapRes := make(map[string]interface{})
	err := gb.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
	fmt.Println("userMapper.findMapById-->", mapRes, err)
	fmt.Println("############################ --- example2")
}

func example3() {
	db, _ := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	dbs := make(map[string]*gobatis.GoBatisDB)
	dbs["ds1"] = gobatis.NewGoBatisDB(gobatis.DBTypeMySQL, db)

	option := gobatis.NewDBOption().
		DB(dbs).
		ShowSQL(true).
		Mappers([]string{"mapper/userMapper.xml"})

	gobatis.Init(option)

	gb := gobatis.Get("ds1")

	mapRes := make(map[string]interface{})
	err := gb.Select("userMapper.findMapById", map[string]interface{}{"id": 1})(mapRes)
	fmt.Println("userMapper.findMapById-->", mapRes, err)
	fmt.Println("############################ --- example3")
}
