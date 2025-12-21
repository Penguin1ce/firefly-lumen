package handler

import (
	"firflybot/config"

	"github.com/go-telegram/bot"
)

// 在包级别缓存 config，方便其他 handler 使用。
var cfg *config.Config

// RegisterHandlers 统一注册所有 bot handler。
func RegisterHandlers(b *bot.Bot, c *config.Config) {
	if b == nil || c == nil {
		return
	}

	cfg = c

	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, AIHandler)
}
