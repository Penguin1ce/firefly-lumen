package aihelper

import (
	"fireflybot/db"
	"fireflybot/model"
	"sync"

	"github.com/sirupsen/logrus"
)

type AIHelper struct {
	Client    model.OpenAIClient
	mu        sync.RWMutex
	messages  []*db.History
	SessionId string
}

func NewAIHelper(client model.OpenAIClient, sessionId string) *AIHelper {
	return &AIHelper{
		Client:    client,
		SessionId: sessionId,
		messages:  make([]*db.History, 0),
	}
}

func (ah *AIHelper) AddMessage(message string, sessionId string, isUser bool) {
	history := &db.History{
		SID:     sessionId,
		Message: message,
		IsUser:  isUser,
	}
	ah.messages = append(ah.messages, history)
	err := db.AppendMessage(sessionId, message, isUser)
	if err != nil {
		return
	}
	logrus.Info("插入消息成功：%s", message)
}
