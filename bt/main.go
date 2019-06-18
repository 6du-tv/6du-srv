package main

import (
	util "bt/util"

	. "github.com/urwork/throw"

	"fmt"
	"log"
	"net"
	"time"
)

type CMD uint8

const (
	ALIVE CMD = iota
	NODE
)

/*
发送命令
回复命令

发送数据 数据hash 当前是第几个包 数据有多少个包 数据
接受数据 数据hash

*/

func main() {
	print(util.CONFIG.ID)
	localAddr, err := net.ResolveUDPAddr("udp", ":20000")
	Throw(err)

	remoteAddr, err := net.ResolveUDPAddr("udp", "47.105.53.166:20000")
	Throw(err)

	// Build listening connections
	conn, err := net.ListenUDP("udp", localAddr)
	// Exit if some error occured
	defer conn.Close()
	Throw(err)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		// write a message to server
		for range ticker.C {
			_, err = conn.WriteToUDP([]byte(util.B64uuid()), remoteAddr)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(">>> Packet sent to: ", remoteAddr)
			}
		}
	}()

	for {
		// Receive response from server
		buf := make([]byte, util.CONFIG.MTU)
		rn, remAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, remAddr, string(buf[:rn]))
		}
	}
}
