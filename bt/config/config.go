package config

import (
	"bytes"
	"crypto/ecdsa"
	"io/ioutil"
	"math/rand"
	"os"
	"path"

	udp "bt/udp"
	cryptoRand "crypto/rand"

	"github.com/btcsuite/btcd/btcec"

	"github.com/BurntSushi/toml"
	. "github.com/urwork/throw"
)

type Path struct {
	KEY string
}

type Net struct {
	MTU  uint16
	PORT uint16
}

type Config struct {
	PATH Path
	NET  Net
}

var SECRET []byte

var CONFIG Config

var ROOT = Root()

func init() {
	initConfig()
	initKey()
}

func initKey() {
	filepath := path.Join(ROOT, CONFIG.PATH.KEY)
	var err error
	SECRET, err = ioutil.ReadFile(filepath)

	if os.IsNotExist(err) || (err == nil && len(SECRET) == 0) {

		// 参考 https://github.com/ethereum/go-ethereum/blob/master/crypto/crypto.go
		// privateKey := hex.EncodeToString()
		// address := crypto.PubkeyToAddress(key.PublicKey).Hex()

		key, _ := ecdsa.GenerateKey(btcec.S256(), cryptoRand.Reader)

		SECRET = key.D.Bytes()

		Throw(ioutil.WriteFile(filepath, SECRET, 0400))

	}
	print("SECRET LEN ", len(SECRET))
}

func initConfig() {
	filepath := path.Join(ROOT, "config.toml")

	_, err := os.Stat(filepath)
	if !os.IsNotExist(err) {
		_, err := toml.DecodeFile(filepath, &CONFIG)
		Throw(err)
	}

	update := false

	if 0 == len(CONFIG.PATH.KEY) {
		key, err := os.Hostname()
		if err != nil {
			key = "default"
		}

		CONFIG.PATH.KEY = key + ".key"
	}

	if 0 == CONFIG.NET.MTU {
		CONFIG.NET.MTU = 1472
		update = true
	}

	if 0 == CONFIG.NET.PORT {
		port := uint16(rand.Int31n(20000)) + 10000
		for ; port < 49000; port++ {
			if !udp.PortUsed(port) {
				break
			}
		}

		CONFIG.NET.PORT = port
		update = true
	}

	if update {
		b := &bytes.Buffer{}
		Throw(toml.NewEncoder(b).Encode(CONFIG))

		Throw(ioutil.WriteFile(filepath, b.Bytes(), os.ModePerm))
	}
}
