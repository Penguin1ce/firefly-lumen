package main

import (
	"context"
	"fmt"
	"log"

	myconfig "fireflybot/config"
	"fireflybot/handler"
	"fireflybot/sql"

	"github.com/go-telegram/bot"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := myconfig.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, err := bot.New(config.TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	// 统一注册所有的 handler
	handler.RegisterHandlers(b, config)

	user := "root"
	pass := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "firefly"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接错误")
	}
	if err := database.AutoMigrate(&sql.Session{}); err != nil {
		log.Fatal(err)
	}
	log.Println("流萤 Bot 成功运行")
	b.Start(context.TODO())
}
