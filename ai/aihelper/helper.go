package aihelper

import (
	"fireflybot/db"
	"fireflybot/model"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type AIHelper struct {
	Client    *model.OpenAIClient
	mu        sync.RWMutex
	messages  []*db.History
	SessionId string
}

func NewAIHelper(client *model.OpenAIClient, sessionId string) *AIHelper {
	return &AIHelper{
		Client:    client,
		SessionId: sessionId,
		messages:  make([]*db.History, 0, 20),
	}
}

func (ah *AIHelper) AddMessage(message string, isUser bool) error {
	sessionId := ah.SessionId
	err := db.AppendMessage(sessionId, message, isUser)
	if err != nil {
		logrus.WithError(err).Error("插入消息到数据库失败")
		return err
	}
	history := &db.History{
		SID:     sessionId,
		Message: message,
		IsUser:  isUser,
	}
	ah.mu.Lock()
	ah.messages = append(ah.messages, history)
	ah.mu.Unlock()
	logrus.Infof("插入消息成功：%s", message)
	return nil
}

func (ah *AIHelper) GetAllMessageFromHistory() ([]*db.History, error) {
	sessionId := ah.SessionId
	// GetAllMessagesOrderByTime返回的是[]history
	history, err := db.GetAllMessagesOrderByTime(sessionId)
	if err != nil {
		return nil, fmt.Errorf("获取全部历史记录错误：%v", err)
	}
	logrus.Info("获得全部历史记录成功： " + sessionId)
	ah.mu.Lock()
	ah.messages = history
	ah.mu.Unlock()
	return history, nil
}

func (ah *AIHelper) DeleteAllHistory() error {
	sid := ah.SessionId
	// 1. 先删除数据库的
	if err := db.DeleteAllMessages(sid); err != nil {
		logrus.WithError(err).Error("删除数据库历史消息失败")
		return err
	}

	ah.mu.Lock()
	ah.messages = make([]*db.History, 0, 20)
	ah.mu.Unlock()

	logrus.Info("删除历史记录成功： " + sid)
	return nil
}
