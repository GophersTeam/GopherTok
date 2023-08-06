package test

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite" // 使用SQLite作为测试数据库
	"gorm.io/gorm"
	"testing"
)

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

func TestUpdateUserCounts(t *testing.T) {
	// 创建测试数据库连接
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	defer db.Migrator().DropTable(&User{}) // 在测试结束后删除表

	// 自动迁移表结构
	err = db.AutoMigrate(&User{})
	assert.NoError(t, err)

	// 插入测试数据
	user := User{
		ID:            1,
		FollowCount:   5,
		FollowerCount: 10,
		FriendCount:   15,
	}
	err = db.Create(&user).Error
	assert.NoError(t, err)

	// 要更新的计数值
	newFollowCount := int64(8)
	newFollowerCount := int64(12)
	newFriendCount := int64(18)

	// 更新用户计数
	err = UpdateUserCounts(db, user.ID, newFollowCount, newFollowerCount, newFriendCount)
	assert.NoError(t, err)

	// 查询更新后的数据
	var updatedUser User
	err = db.First(&updatedUser, user.ID).Error
	assert.NoError(t, err)

	// 验证更新结果
	assert.Equal(t, newFollowCount, updatedUser.FollowCount)
	assert.Equal(t, newFollowerCount, updatedUser.FollowerCount)
	assert.Equal(t, newFriendCount, updatedUser.FriendCount)
}

func TestMain(m *testing.M) {
	// 运行测试前的初始化操作
	m.Run()
}
func UpdateUserCounts(db *gorm.DB, userID int64, followCount, followerCount, friendCount int64) error {
	return db.Model(&User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"follow_count":  followCount,
		"FollowerCount": followerCount,
		"FriendCount":   friendCount,
	}).Error
}
