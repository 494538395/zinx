package net

import (
	"bytes"
	"encoding/binary"
	"github.com/494538395/zinx/iface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (p *DataPack) GetHeadLen() uint32 {
	// head:  datalen | msgID
	//        uint32 4   uint32 4
	// 两个int，即 8 个字节
	return 8
}

// Pack 封包
func (p *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	// 创建一个存放 byte[] 的缓冲
	buffer := bytes.NewBuffer([]byte{})

	// 将 datalen 写入
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	// 将 msgID 写入
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	// 将 data 写入
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// UnPack 拆包
func (p *DataPack) UnPack(data []byte) (iface.IMessage, error) {
	// 读取 head 信息，然后根据 head 里的 datalen 读取数据内容
	msg := &Message{}

	reader := bytes.NewReader(data)

	// 读取 datalen
	if err := binary.Read(reader, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读取 msgID
	if err := binary.Read(reader, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	return msg, nil
}
