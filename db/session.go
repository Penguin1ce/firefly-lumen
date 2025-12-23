package db

import "time"

// Session 对应Telegram的对话，个人对话一人一个chatId,用chatId作为主键
type Session struct {
	ID        uint   `gorm:"primaryKey"`
	SID       string `gorm:"column:s_id;type:varchar(64);not null;uniqueIndex"`
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func IsActive(sid string) bool {
	session := Session{}
	err := DB.Where("sid = ?", sid).First(&session).Error
	if err != nil {
		return false
	}
	return session.IsActive
}

// CreateSession 如果该用户的chatId之前没有激活，则插入该用户的chatId
func CreateSession(sid string) (*Session, error) {
	session := &Session{}
	err := DB.
		Where("s_id = ?", sid).
		FirstOrCreate(session, Session{
			SID:      sid,
			IsActive: true,
		}).Error
	return session, err
}

// ExistsSession 查询该用户是否存在
func ExistsSession(sid string) bool {
	session := Session{}
	err := DB.Where("sid = ?", sid).First(&session).Error
	return err == nil
}
