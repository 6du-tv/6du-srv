package net

import (
	"fmt"
	"net"

	. "github.com/urwork/throw"
)

func PortUsed(port uint16) bool {
	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	Throw(err)

	conn, err := net.ListenUDP("udp", localAddr)

	if err != nil {
		return true
	} else {
		defer conn.Close()
		return false
	}
}
