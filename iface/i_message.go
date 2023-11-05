package iface

// IMessage 消息接口
type IMessage interface {
	GetMsgID() uint32   // 获取消息 ID
	GetMsgLen() uint32  // 获取消息长度
	GetMsgData() []byte // 获取消息内容
	SetMsgID(uint32)    // 设置消息 ID
	SetMsgData([]byte)  // 设置消息内容
	SetMsgLen(uint32)   // 设置消息长度
}
