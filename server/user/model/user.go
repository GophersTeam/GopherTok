package model

// User 定义结构体映射表格
type User struct {
	ID              int64  `gorm:"primaryKey;autoIncrement;column:id"`
	Username        string `gorm:"size:32;not null;column:username"`
	Password        string `gorm:"size:32;not null;column:password"`
	Avatar          string `gorm:"size:255;column:avatar"`
	BackgroundImage string `gorm:"size:255;column:background_image"`
	Signature       string `gorm:"size:255;column:signature"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	FriendCount     int64  `gorm:"column:friend_count"`
}
