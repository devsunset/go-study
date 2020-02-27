package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println(&product)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)

	/*
		https://gorm.io/docs
		https://gorm.io/ko_KR/

		https://gorm.io/docs/create.html
		https://gorm.io/docs/query.html
		https://gorm.io/docs/update.html
		https://gorm.io/docs/delete.html
	*/
}
