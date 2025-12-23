package db

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type History struct {
	ID        uint   `gorm:"primaryKey"`
	SID       string `gorm:"index"`
	Message   string
	IsUser    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllMessagesOrderByTime(sid string) ([]*History, error) {
	var history []*History
	err := DB.Where("sid = ?", sid).Order("created_at ASC").Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

// AppendMessage 用户消息和AI聊天记录插入到数据库
func AppendMessage(sid string, message string, isUser bool) error {
	var history History
	history.SID = sid
	history.Message = message
	history.IsUser = isUser
	err := DB.Create(&history).Error
	if err != nil {
		return fmt.Errorf("插入到历史消息的时候出错：%v", err)
	}
	return nil
}

func DeleteAllMessages(sid string) error {
	err := DB.Where("sid = ?", sid).Delete(&History{}).Error
	if err != nil {
		return err
	}
	logrus.Info("成功删除所有 " + sid + " 的消息")
	return nil
}
