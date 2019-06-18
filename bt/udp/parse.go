package udp

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/scylladb/go-set"
)

const PING_RATE_LIMIT = 1024

var REPLYD = set.NewUint64Set()

func init() {

	ticker := time.NewTicker(60 * time.Second)

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
		ip := binary.BigEndian.Uint64(remote.IP)
		if !REPLYD.Has(ip) && REPLYD.Size() < PING_RATE_LIMIT {
			REPLYD.Add(ip)
			conn.WriteUDP([]byte{byte(PONG)}, remote)
		} else {
			println("remote.IP", remote.IP.String())
			println("IP", ip)
			println("REPLYD.Cardinality()", REPLYD.Size())
			println("REPLYD.Contains(remote.IP)", REPLYD.Has(ip))
			println("REPLYD", REPLYD.String())
		}

	default:
		fmt.Printf("<<<  %d bytes received from: %v, data: %x\n", len(buf), remote, buf)

	}
}
