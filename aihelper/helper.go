package aihelper

import (
	"fireflybot/model"
	"sync"

	"github.com/cloudwego/eino/schema"
)

type AIHelper struct {
	Client   model.OpenAIClient
	mu       sync.RWMutex
	messages []*schema.Message
}
