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
		ip := []byte(remote.IP)
		if !REPLYD.Contains(ip) && REPLYD.Cardinality() < 1024 {
			REPLYD.Add(ip)
			conn.WriteUDP([]byte{byte(PONG)}, remote)
		} else {
			print("IP", ip)
			print("REPLYD.Cardinality()", REPLYD.Cardinality())
			print("REPLYD.Contains(remote.IP)", REPLYD.Contains(ip))
			print("REPLYD", REPLYD)
		}

	default:
		fmt.Printf("<<<  %d bytes received from: %v, data: %x\n", len(buf), remote, buf)

	}
}
