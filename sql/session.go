package sql

import "time"

type Session struct {
	ID       uint   `gorm:"primaryKey"`
	SID      string `gorm:"index"`
	CreateAt time.Time
	UpdateAt time.Time
}
