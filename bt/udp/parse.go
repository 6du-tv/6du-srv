package udp

import (
	"fmt"
	"net"
)

func Parse(buf []byte, remote *net.UDPAddr, conn *Conn) {
	cmd := CMD(buf[0])

	switch cmd {

	case PING:
		conn.WriteUDP([]byte{byte(PONG)}, remote)
	default:
		fmt.Printf("<<<  %d bytes received from: %v, data: %x\n", len(buf), remote, buf)

	}
}
