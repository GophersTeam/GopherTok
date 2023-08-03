package model

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID         int64          `gorm:"primaryKey"`
	UserID     int64          `gorm:"column:user_id;not null"`
	Title      string         `gorm:"column:title;not null"`
	PlayURL    string         `gorm:"column:play_url;not null"`
	CoverURL   string         `gorm:"column:cover_url;not null"`
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}
