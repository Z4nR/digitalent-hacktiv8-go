package config

import (
	"go-first/model"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zulham4!21N0rR12!dh@tcp(127.0.0.1:3306)/mylocal?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(model.Order{})
	db.AutoMigrate(model.Item{})
	db.Model(model.Item{}).AddForeignKey("order_id", "orders(order_id)", "CASCADE", "CASCADE")
	return db
}
