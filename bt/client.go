package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var raddr = flag.String("raddr", "xvc.bid:20000", "remote server address")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	// Resolving Address
	remoteAddr, err := net.ResolveUDPAddr("udp", *raddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Make a connection
	tmpAddr := &net.UDPAddr{
		IP:   nil,
		Port: 0,
	}

	conn, err := net.DialUDP("udp", tmpAddr, remoteAddr)
	// Exit if some error occured
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	// write a message to server
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(">>> Packet sent to: ", *raddr)
	}

	// Receive response from server
	buf := make([]byte, 1024)
	rn, rmAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, rmAddr, string(buf[:rn]))
	}
}
