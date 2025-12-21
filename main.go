package main

import (
	"context"
	myconfig "firflybot/config"
	"firflybot/handler"
	"log"

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

	log.Println("流萤 Bot 成功运行")
	b.Start(context.TODO())
}
