package sql

import "time"

type History struct {
	ID       uint   `gorm:"primaryKey"`
	SID      string `gorm:"index"`
	Message  string
	CreateAt time.Time
	UpdateAt time.Time
}
