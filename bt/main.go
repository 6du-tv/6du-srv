package main

import (
	util "bt/util"

	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	. "github.com/urwork/throw"
)

type Config struct {
	ID string
}

var CONFIG Config

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(ok)
	}

	dirname := path.Dir(filename)
	filepath := path.Join(dirname, "config.toml")

	if _, err := os.Stat(filepath); os.IsNotExist(err) {

	} else {

		_, err := toml.DecodeFile(filepath, &CONFIG)
		Throw(err)

	}

	if 0 == len(CONFIG.ID) {
		CONFIG.ID = util.B64uuid()
		b := &bytes.Buffer{}
		encoder := toml.NewEncoder(b)

		err := encoder.Encode(CONFIG)
		Throw(err)

		Throw(ioutil.WriteFile(filepath, b.Bytes(), 0644))

	}

	print(CONFIG.ID)

}

const MTU int = 1472

type CMD uint8

const (
	ALIVE CMD = iota
	NODE
)

/*
发送命令
回复命令

发送数据 数据hash 当前是第几个包 数据有多少个包 数据
接受数据 数据hash

*/

func main() {

	localAddr, err := net.ResolveUDPAddr("udp", ":20000")
	Throw(err)

	remoteAddr, err := net.ResolveUDPAddr("udp", "47.105.53.166:20000")
	Throw(err)

	// Build listening connections
	conn, err := net.ListenUDP("udp", localAddr)
	// Exit if some error occured
	defer conn.Close()
	Throw(err)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		// write a message to server
		for range ticker.C {
			_, err = conn.WriteToUDP([]byte(util.B64uuid()), remoteAddr)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(">>> Packet sent to: ", remoteAddr)
			}
		}
	}()

	for {
		// Receive response from server
		buf := make([]byte, MTU)
		rn, remAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, remAddr, string(buf[:rn]))
		}
	}
}
