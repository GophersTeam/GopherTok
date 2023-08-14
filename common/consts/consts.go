package consts

const (
	UserId = "userId"

	MsgTypeRecv = 0
	MsgTypeSend = 1
	CommentAdd  = 1
	CommentDel  = 2
)

const (
	UserMachineId = iota
	ChatMachineId
	VideoMachineId
)

const (
	Token = "token"
)

// StoreType 存储类型(表示文件存到哪里)
type StoreType int

const (
	_ StoreType = iota
	// StoreLocal : 节点本地
	StoreLocal
	// StoreMinio : Minio集群
	StoreMinio
	// StoreCOS : 腾讯云COS
	StoreCOS
)

const (
	MinioBucket = "gophertok-video"
)
const (
	CoverTemp = "/home/project/gophertok/temp/"
)
