package multiCli

import (
	"fmt"
	"net"
	"testing"
)

func TestCli01(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 512)
		readCnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("read cnt-->", readCnt)
		fmt.Println("read content-->", string(buf))
	}
}

func TestCli02(t *testing.T) {

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 512)
		readCnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("read cnt-->", readCnt)
		fmt.Println("read content-->", string(buf))
	}
}

func TestCli03(t *testing.T) {

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 512)
		readCnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("read cnt-->", readCnt)
		fmt.Println("read content-->", string(buf))
	}

}
