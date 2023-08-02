// Code generated by goctl. DO NOT EDIT.
package types

type Res struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavourited string `json:"total_favourited"`
	WorkCount       int64  `json:"work_count"`
	FavouriteCount  int64  `json:"favourite_count"`
}

type Follow struct {
	Id         int64 `gorm:"id" json:"id"`
	UserId     int64 `gorm:"user_id" json:"user_id"`
	FollowerId int64 `gorm:"fowwower_id" json:"follower_id"`
	IsFollow   bool  `gorm:"is_follow" json:"is_follow"`
}

type FollowReq struct {
	Token      string `form:"token"`
	ToUserId   string `form:"to_user_id"`
	ActionType string `form:"action_type"`
}

type FollowRes struct {
	Res
}

type FollowListReq struct {
	UserId string `form:"user_id"`
	Token  string `form:"token"`
}

type FollowListRes struct {
	Res
	UserList []User `json:"user_list"`
}

type FollowerListReq struct {
	UserId string `form:"user_id"`
	Token  string `form:"token"`
}

type FollowerListRes struct {
	Res
	UserList []User `json:"user_list"`
}

type FriendListReq struct {
	UserId string `form:"user_id"`
	Token  string `form:"token"`
}

type FriendListRes struct {
	Res
	UserList []User `json:"user_list"`
}
