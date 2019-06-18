package config

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"runtime"

	udp "bt/udp"
	util "bt/util"

	. "github.com/urwork/throw"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SECRET string
	MTU    uint16
	PORT   uint16
}

var CONFIG Config

func init() {
	_, filename, _, _ := runtime.Caller(0)

	dirname := path.Dir(filename)
	filepath := path.Join(dirname, "config.toml")

	_, err := os.Stat(filepath)
	if !os.IsNotExist(err) {
		_, err := toml.DecodeFile(filepath, &CONFIG)
		Throw(err)
	}

	update := false

	if 0 == len(CONFIG.SECRET) {
		CONFIG.SECRET = util.RandByteB64(32)
		update = true
	}

	if 0 == CONFIG.MTU {
		CONFIG.MTU = 1472
		update = true
	}

	if 0 == CONFIG.PORT {
		port := uint16(rand.Int31n(20000)) + 10000
		for ; port < 49000; port++ {
			if !udp.PortUsed(port) {
				break
			}
		}

		CONFIG.PORT = port
		update = true
	}

	if update {
		b := &bytes.Buffer{}
		Throw(toml.NewEncoder(b).Encode(CONFIG))
		Throw(ioutil.WriteFile(filepath, b.Bytes(), 0644))
	}
}
