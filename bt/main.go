package main

import (
	util "bt/util"
	"fmt"
	"log"
	"net"
	"time"
)

func init() {
}

const MTU int = 1472

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

	localAddr, err := net.ResolveUDPAddr("udp", ":20000")
	if err != nil {
		fmt.Println(err)
		return
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", "47.105.53.166:20000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Build listening connections
	conn, err := net.ListenUDP("udp", localAddr)
	// Exit if some error occured
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

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
		buf := make([]byte, MTU)
		rn, remAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, remAddr, string(buf[:rn]))
		}
	}
}
