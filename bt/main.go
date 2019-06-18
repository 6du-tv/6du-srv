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

type UDPConn struct {
	*net.UDPConn
}

func (conn *UDPConn) WriteUDP(b []byte, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP(b, addr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(">>> Packet sent to: ", addr)
	}
}

func main() {
	print(CONFIG.ID)
	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", CONFIG.PORT))
	Throw(err)

	remoteAddr, err := net.ResolveUDPAddr("udp", "47.105.53.166:27039")
	Throw(err)

	// Build listening connections
	_conn, err := net.ListenUDP("udp", localAddr)
	Throw(err)
	defer _conn.Close()
	conn := &UDPConn{_conn}

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		// write a message to server
		for range ticker.C {
			buffer := []byte{byte(udp.PING)}
			//		append(buffer, []byte(util.B64uuid()))
			conn.WriteUDP(buffer, remoteAddr)

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

func Parse(buf []byte, remote *net.UDPAddr, conn *UDPConn) {
	fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", len(buf), remote, buf)
	cmd := udp.CMD(buf[0])

	switch cmd {

	case udp.PING:
		conn.WriteUDP([]byte{byte(udp.PONG)}, remote)
	default:
		fmt.Printf("<<<  %d bytes received from: %v, data: %x\n", len(buf), remote, buf)

	}
}
