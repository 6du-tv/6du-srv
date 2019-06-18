package main

import (
	. "bt/config"
	udp "bt/udp"

	. "github.com/urwork/throw"

	"fmt"
	"log"
	"net"
	"time"
)

/*
发送命令
回复命令

发送数据 数据hash 当前是第几个包 数据有多少个包 数据
接受数据 数据hash

*/

func main() {
	print(CONFIG.ID)
	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", CONFIG.PORT))
	Throw(err)

	remoteAddr, err := net.ResolveUDPAddr("udp", "47.105.53.166:27039")
	Throw(err)

	// Build listening connections
	conn, err := net.ListenUDP("udp", localAddr)
	Throw(err)
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		// write a message to server
		for range ticker.C {
			buffer := []byte{byte(udp.PING)}
			//		append(buffer, []byte(util.B64uuid()))
			_, err = conn.WriteToUDP(buffer, remoteAddr)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(">>> Packet sent to: ", remoteAddr)
			}
		}
	}()

	for {
		// 根据UDP协议，从UDP数据包的包头可以看出，UDP的最大包长度是2^16-1的个字节
		// 由于UDP包头占8个字节，而在IP层进行封装后的IP包头占去20字节，所以这个是UDP数据包的最大理论长度是2^16 - 1 - 8 - 20 = 65507字节
		buf := make([]byte, 65507)

		rn, remAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		} else {
			Parse(buf[:rn], remAddr, conn)
		}
	}
}

func Parse(buf []byte, remote *net.UDPAddr, conn *net.UDPConn) {
	fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", len(buf), remote, buf)
	cmd := udp.CMD(buf[0])

	switch cmd {

	case udp.PING:
		conn.WriteToUDP([]byte{byte(udp.PONG)}, remote)
	default:
		fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", len(buf), remote, buf)

	}
}
