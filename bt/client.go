package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	localAddr, err := net.ResolveUDPAddr("udp", ":20000")
	if err != nil {
		fmt.Println(err)
		return
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", "xvc.bid:20000")
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

	// write a message to server
	_, err = conn.WriteToUDP([]byte("hello"), remoteAddr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(">>> Packet sent to: ", remoteAddr)
	}

	// Receive response from server
	buf := make([]byte, 1024)
	rn, remAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, remAddr, string(buf[:rn]))
	}
}
