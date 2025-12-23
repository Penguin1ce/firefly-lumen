package model

import (
	"context"

	"fireflybot/db"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type AIModel interface {
	GenerateResponse(ctx context.Context, messages []*schema.Message) (*schema.Message, error)
	CreateTemplate() prompt.ChatTemplate
	CreateMessagesFromTemplate(text string, history []*db.History) []*schema.Message
}
