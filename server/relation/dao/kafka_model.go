package dao

type CountData struct {
	FollowCountKey string `json:"follow_count_key"`
	FollowCount    string `json:"follow_count"`

	FollowerCountKey string `json:"follower_count_key"`
	FollowerCount    string `json:"follower_count"`

	FriendCountKey string `json:"friend_count_key"`
	FriendCount    string `json:"friend_count"`
}

type FollowData struct {
	Method string `json:"method"`

	Id         int64 `gorm:"id" json:"id"`
	UserId     int64 `gorm:"user_id" json:"user_id"`
	FollowerId int64 `gorm:"follower_id" json:"follower_id"`
	IsFollow   bool  `gorm:"is_follow" json:"isFollow"`
}
