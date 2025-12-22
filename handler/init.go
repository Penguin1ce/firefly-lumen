package handler

import (
	"fireflybot/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// 在包级别缓存 config，方便其他 handler 使用。
var cfg *config.Config

// RegisterHandlers 统一注册所有 bot handler。
func RegisterHandlers(b *bot.Bot, c *config.Config) {
	if b == nil || c == nil {
		return
	}

	cfg = c
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypePrefix, HelpHandler)
	// 捕获贴纸
	b.RegisterHandlerMatchFunc(
		func(u *models.Update) bool {
			return u.Message != nil && u.Message.Sticker != nil
		},
		StickerHandler,
	)
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, AIHandler)
}
