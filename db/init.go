package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	user := "root"
	pass := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "firefly"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = database
	if err != nil {
		log.Fatal("数据库连接错误")
		return err
	}
	if err := database.AutoMigrate(&Session{}, &History{}); err != nil {
		log.Fatal(err)
		return err
	}
	DB = database
	return nil
}
