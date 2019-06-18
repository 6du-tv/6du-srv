package udp

import (
	"fmt"
	"net"
	"time"

	mapset "github.com/deckarep/golang-set"
)

var REPLYD = mapset.NewSet()

func init() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		// write a message to server
		for range ticker.C {
			REPLYD.Clear()
		}
	}()
}

func Parse(buf []byte, remote *net.UDPAddr, conn *Conn) {
	cmd := CMD(buf[0])

	switch cmd {

	case PING:
		if !REPLYD.Contains(remote.IP) && REPLYD.Cardinality() < 1024 {
			REPLYD.Add(IP)
			conn.WriteUDP([]byte{byte(PONG)}, remote)
		} else {
			print("REPLYD.Cardinality()", REPLYD.Cardinality())
			print("REPLYD.Contains(remote.IP)", REPLYD.Contains(remote.IP))
			print("REPLYD", REPLYD)
		}

	default:
		fmt.Printf("<<<  %d bytes received from: %v, data: %x\n", len(buf), remote, buf)

	}
}
