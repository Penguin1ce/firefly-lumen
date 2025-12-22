package handler

import (
	"context"
	"log"
	"math/rand"
	"time"

	"fireflybot/service"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// AIHandler 它会把用户输入转发给 AI Service，并将回复再发回用户。
func AIHandler(ctx context.Context, b *bot.Bot, upd *models.Update) {
	if upd == nil || upd.Message == nil {
		return
	}

	reply := service.AiChatService(ctx, upd.Message.Text)
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
	
	_ = SendRandomSticker(ctx, b, upd.Message.Chat.ID)

}

var randomStickers = []string{
	"CAACAgUAAxkBAAIB5WlIE3f1_LyOwrOb4qn64NPTx6lLAAJ5EwACokoBVr_syLyqIg9iNgQ",
	"CAACAgUAAxkBAAIB4GlIEqhu15xl9Sq0CwKX7ShGUOHYAAIdEQACSAoAAVZwXoQGZWI-_TYE",
	"CAACAgEAAxkBAAIB4mlIE0shoyWjRMTDrpKN7PrYVpofAAMDAAKuPrBEU7bvhD4nBrs2BA",
	"CAACAgUAAxkBAAIB6GlIE8MO7CoOrrzd15f4MVbUcGtnAAKbEwACy2wBVm-9lB7VjI0sNgQ",
	"CAACAgUAAxkBAAIB7mlIE9n_Uq2vHJ9NkUfE7Vq1VSYFAAKoEwACDYQAAVZMk-IZYs-PwDYE",
}

func SendRandomSticker(
	ctx context.Context,
	b *bot.Bot,
	chatID int64,
) error {

	rand.Seed(time.Now().UnixNano())
	stickerID := randomStickers[rand.Intn(len(randomStickers))]

	_, err := b.SendSticker(ctx, &bot.SendStickerParams{
		ChatID: chatID,
		Sticker: &models.InputFileString{
			Data: stickerID,
		},
	})

	return err
}
