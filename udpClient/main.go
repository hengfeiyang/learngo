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
	// dial udp server
	ipaddr := "127.0.0.1"
	port := 8080
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.ParseIP(ipaddr),
		Port: port,
	})
	if err != nil {
		fmt.Println("connect error!", err)
		return
	}
	socket.SetWriteDeadline(time.Now().Add(2 * time.Second))
	socket.SetReadDeadline(time.Now().Add(2 * time.Second))
	defer socket.Close()

	// 发送数据
	_, err = socket.Write(msg)
	if err != nil {
		fmt.Println("send data error!", err)
		return
	}

	// 接收数据
	data := bytes.NewBuffer(nil)
	var buf [128]byte
	n, err := socket.Read(buf[0:])
	data.Write(buf[0:n])
	if err != nil {
		fmt.Println("read data error!", err)
	}
	result := data.Bytes()
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
