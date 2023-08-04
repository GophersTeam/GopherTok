package model

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID          int64          `gorm:"primaryKey"`
	UserID      int64          `gorm:"column:user_id;not null"`
	Title       string         `gorm:"column:title;not null;size:255"`
	PlayURL     string         `gorm:"column:play_url;not null;size:255"`
	CoverURL    string         `gorm:"column:cover_url;not null;size:255"`
	CreateTime  time.Time      `gorm:"column:create_time"`
	UpdateTime  time.Time      `gorm:"column:update_time"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time"`
	VideoSha256 string         `gorm:"column:video_sha256;size:255"`
}
