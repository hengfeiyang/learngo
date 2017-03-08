package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

func send(msg []byte) {
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(104, 207, 150, 187),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	socket.SetDeadline(time.Now().Add(2 * time.Second))
	defer socket.Close()

	// 发送数据
	//senddata := []byte("hello server!")
	_, err = socket.Write(msg)
	if err != nil {
		fmt.Println("发送数据失败!", err)
		return
	}

	// 接收数据
	data := bytes.NewBuffer(nil)
	var buf [128]byte
	var remoteAddr interface{}
	for {
		n, addr, err := socket.ReadFromUDP(buf[0:])
		data.Write(buf[0:n])
		if err != nil {
			fmt.Println("读取数据失败!", err)
			break
		}
		remoteAddr = addr
	}
	result := data.Bytes()
	fmt.Println(len(result), remoteAddr)
	fmt.Printf("%s\n", string(result))
}

func checkNetOpError(err error) error {
	if err != nil {
		netOpError, ok := err.(*net.OpError)
		if ok && netOpError.Err.Error() == "use of closed network connection" {
			return nil
		}
	}
	return err
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command:")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "quit" || line == "exit" {
			break
		}
		send([]byte(line))
	}
}
