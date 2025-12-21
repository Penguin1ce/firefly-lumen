package model

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type AIModel interface {
	GenerateResponse(ctx context.Context, messages []*schema.Message) (*schema.Message, error)
	CreateTemplate() prompt.ChatTemplate
	CreateMessagesFromTemplate(text string) []*schema.Message
}
