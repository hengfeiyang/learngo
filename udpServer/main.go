package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// ListenUDP
	ipaddr := "0.0.0.0"
	port := 8080
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.ParseIP(ipaddr),
		Port: port,
	})
	if err != nil {
		log.Println("Listen error!", err)
		return
	}
	log.Println("Listen on:", port)

	defer socket.Close()

	for {
		// read data
		data := make([]byte, 100)
		size, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Println("read data error!", err)
			continue
		}
		msg := fmt.Sprintf("from [%s] received %d bytes: \n%s\n", remoteAddr, size, data)
		log.Print(msg)

		// send data
		_, err = socket.WriteToUDP([]byte(msg), remoteAddr)
		if err != nil {
			log.Println("send data error!", err)
		} else {
			log.Println("reply ok")
		}
	}
}
