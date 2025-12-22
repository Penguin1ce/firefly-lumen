package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, upd *models.Update) {
	if upd == nil || upd.Message == nil {
		return
	}
	chatID := upd.Message.Chat.ID
	text := upd.Message.Text

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   strconv.FormatInt(chatID, 10) + text,
	})
	if err != nil {
		log.Printf("send message error: %v", err)
	}
}

func StickerHandler(ctx context.Context, b *bot.Bot, upd *models.Update) {
	st := upd.Message.Sticker

	// 贴纸的 file_id（用来 sendSticker）
	fileID := st.FileID

	// 贴纸包名（用来 GetStickerSet）
	setName := st.SetName

	text := fmt.Sprintf("收到贴纸 \nfile_id: %s\nset_name: %s", fileID, setName)

	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   text,
	})

}
