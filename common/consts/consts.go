package consts

const (
	UserId         = "userId"
	ChatMachineId  = 3
	UserMachineId  = 4
	VideoMachineId = 5
	MsgTypeRecv    = 0
	MsgTypeSend    = 1
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
	CoverTemp = "/Users/liuxian/temp/"
)
