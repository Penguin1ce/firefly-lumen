package db

import (
	"fmt"
	"time"
)

type History struct {
	ID        uint   `gorm:"primaryKey"`
	SID       string `gorm:"index"`
	Message   string
	IsUser    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllMessagesOrderByTime(sid string) ([]History, error) {
	session := Session{}
	err := DB.Where("sid = ?", sid).First(&session).Error
	if err != nil {
		return nil, err
	}
	var history []History
	err = DB.Where("sid = ?", sid).Order("created_at ASC").Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

// AppendMessage 用户消息和AI聊天记录插入到数据库
func AppendMessage(sid string, message string, isUser bool) error {
	session := Session{}
	err := DB.Where("sid = ?", sid).First(&session).Error
	if err != nil {
		return err
	}
	var history History
	history.SID = sid
	history.Message = message
	history.IsUser = isUser
	err = DB.Create(&history).Error
	if err != nil {
		return fmt.Errorf("插入到历史消息的时候出错：%v", err)
	}
	return nil
}
