package handler

import (
	"context"
	"log"

	"firflybot/service"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// AIHandler 它会把用户输入转发给 AI Service，并将回复再发回用户。
func AIHandler(ctx context.Context, b *bot.Bot, upd *models.Update) {
	if upd == nil || upd.Message == nil {
		return
	}

	reply := service.AiChatService(cfg)
	if reply == "" {
		reply = "暂时无法获取回复，请稍后再试。"
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   reply,
	})
	if err != nil {
		log.Printf("send message error: %v", err)
	}
}
