package main

import (
	"bt/util"
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"path"
	"runtime"
)

func connect(uri string) {
	fmt.Println(uri)

	u, err := url.Parse(uri)
	if err != nil {
		log.Printf("url.Parse error %s\n", err)
		return
	}

	addr, err := net.ResolveUDPAddr("udp", u.Host)
	if err != nil {
		log.Printf("Resolve DNS error, %s\n", err)
		return
	}
	log.Print(addr)
}

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	dirpath := path.Dir(currentFilePath)
	file, err := os.Open(path.Join(dirpath, "udp.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wait util.Wait

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var uri string = scanner.Text()
		if len(uri) > 0 {
			wait.Add(func() { connect(uri) })
		}
	}
	wait.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
