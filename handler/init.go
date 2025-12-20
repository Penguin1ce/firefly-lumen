package handler

import (
	"firflybot/config"

	"github.com/go-telegram/bot"
)

// RegisterHandlers 统一注册所有 bot handler。
func RegisterHandlers(b *bot.Bot, cfg *config.Config) {
	if b == nil || cfg == nil {
		return
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, NewAIHandler(cfg.AiUrl, cfg.AiKey))
}
