package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// TestDataPack 封包、拆包测试
func TestDataPack(t *testing.T) {

	// 1. 创建 socket
	listen, err := net.Listen("tcp", ":9991")
	if err != nil {
		panic(err)
	}
	// 2. 监听客户端发送的数据，将数据进行拆包，解析
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				panic(err)
			}

			go func(conn net.Conn) {
				// 拆包
				dp := NewDataPack()

				for {
					headData := make([]byte, dp.GetHeadLen())

					// 读取 headData
					cnt, err := io.ReadFull(conn, headData)
					if err != nil {
						panic(err)
					}
					fmt.Println("cnt-->", cnt)

					// 根据 headData 拆包
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						panic(err)
					}

					if msgHead.GetMsgLen() > 0 {
						// 有数据
						// 第二次从 conn 中读取数据，根据 haeddata 中的 datalen 进行读取
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msgHead.GetMsgLen())
						// 读取 data
						_, err = io.ReadFull(conn, msg.Data)
						if err != nil {
							panic(err)
						}

						fmt.Printf("receive , msgID:%d ,datalen:%d, data:%v ", msg.ID, msg.DataLen, string(msg.Data))
						fmt.Println("读到了数据")
					}

				}
			}(conn)
		}
	}()

	// 2. 模拟客户端连接，并且封包发送数据
	conn, err := net.Dial("tcp", ":9991")
	if err != nil {
		panic(err)
	}

	// 模拟粘包
	// 封装msg1 和 msg2，粘在一起，然后发送给服务端
	dp := NewDataPack()

	msg01 := &Message{ID: 1, DataLen: 3, Data: []byte{'a', 'b', 'c'}}
	sendMsg01, err := dp.Pack(msg01)
	if err != nil {
		panic(err)
	}

	msg02 := &Message{ID: 2, DataLen: 5, Data: []byte{'h', 'e', 'l', 'l', 'o'}}
	sendMsg02, err := dp.Pack(msg02)
	if err != nil {
		panic(err)
	}

	// 将两个包粘在一起
	sendMsg01 = append(sendMsg01, sendMsg02...)
	conn.Write(sendMsg01)

	select {} // 客户端阻塞
}
