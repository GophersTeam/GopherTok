package model

type Comment struct {
	Id         int64  `bson:"_id,omitempty" json:"id,omitempty"`
	Content    string `bson:"content,omitempty" json:"content,omitempty"`
	VideoId    int64  `bson:"video_id,omitempty" json:"video_id,omitempty"`
	UserId     int64  `bson:"user_id,omitempty" json:"user_id,omitempty"`
	CreateDate string `bson:"create_date,omitempty" json:"create_date,omitempty"`
}
