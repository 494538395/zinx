package net

type Message struct {
	ID      uint32 // 消息 ID
	DataLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func NewMessagePackage(msgID uint32, data []byte) *Message {
	return &Message{
		ID:      msgID,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) SetMsgID(ID uint32) {
	m.ID = ID
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}

func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}
