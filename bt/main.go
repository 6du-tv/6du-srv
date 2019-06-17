package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"path"
	"runtime"
	"sync"
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

	var thread sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var uri string = scanner.Text()
		if len(uri) > 0 {
			thread.Add(1)
			go func(uri string) {
				defer thread.Done()
				connect(uri)
			}(uri)
		}
	}
	thread.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
