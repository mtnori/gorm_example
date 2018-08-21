package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product 製品
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "gorm_sample"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// スキーマのマイグレーション
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	db.First(&product, 1)                   // IDが1の製品を取得する
	db.First(&product, "code = ?", "L1212") // codeがL1212の製品を取得する

	db.Model(&product).Update("Price", 2000) //製品のPriceを2000に更新する

	db.Delete(&product) // 製品を削除する
}
