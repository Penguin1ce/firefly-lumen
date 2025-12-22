package main

import (
	"context"
	"fireflybot/db"
	"log"

	myconfig "fireflybot/config"
	"fireflybot/handler"

	"github.com/go-telegram/bot"
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
	log.Println("开始初始化数据库")
	err = db.InitDB()
	if err != nil {
		log.Fatal("数据库初始化失败")
	}
	log.Println("流萤 Bot 成功运行")
	b.Start(context.TODO())
}
