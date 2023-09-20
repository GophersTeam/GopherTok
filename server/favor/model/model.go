package model

type Info struct {
	UserId  int64
	VideoId int64
}

type FavorInfo struct {
	Info
	Method string
}

type Favor struct {
}
