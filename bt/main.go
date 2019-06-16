package main

import (
	"log"
	"time"

	"gitlab.com/axet/libtorrent"
)

func createTorrentFileExample() {
	t1 := libtorrent.CreateTorrentFile("/Users/axet/Downloads/Prattchet")
	ioutil.WriteFile("./test.torrent", t1, 0644)
}

func downloadMagnetWaitExample() {
	libtorrent.Create()
	t1 := libtorrent.AddMagnet("/tmp", "magnet:?...")
	libtorrent.StartTorrent(t1)
	libtorrent.WaitAll()
	log.Println("done")
	libtorrent.Close()
}

func downloadMagnetStatusExample() {
	libtorrent.Create()
	t1 := libtorrent.AddMagnet("/tmp", "magnet:?...")
	libtorrent.StartTorrent(t1)
	for libtorrent.TorrentStatus(t1) == libtorrent.StatusDownloading {
		time.Sleep(100 * time.Millisecond)
		log.Println("loop")
	}
	log.Println("done")
	libtorrent.Close()
}

