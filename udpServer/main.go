package main

import (
	"fmt"
	"net"
)

func main() {
	var s []byte
	s = make([]byte, 0, 2048)
	t := "hahahaha--"
	for i := 0; i < 1000000; i++ {
		s = append(s, t...)
		if len(s) > 1024 {
			break
		}
	}
	fmt.Println(len(s))
	// 创建监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer socket.Close()

	for {
		// 读取数据
		data := make([]byte, 100)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n\n", data)

		// 发送数据
		//senddata := []byte("hello client!")
		senddata := s
		_, err = socket.WriteToUDP(senddata, remoteAddr)
		if err != nil {
			return
			fmt.Println("发送数据失败!", err)
		}
	}
}
