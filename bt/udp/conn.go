package udp

import (
	"fmt"
	"log"
	"net"
)

type Conn struct {
	*net.UDPConn
}

func (conn *Conn) WriteUDP(b []byte, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP(b, addr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf(">>> Packet sent to %s : %x\n", addr, b)
	}
}
