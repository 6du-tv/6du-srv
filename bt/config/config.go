package config

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"runtime"

	"bt/net"
	util "bt/util"

	. "github.com/urwork/throw"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ID   string
	MTU  uint16
	PORT uint16
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

	if 0 == len(CONFIG.ID) {
		CONFIG.ID = util.B64uuid()
		update = true
	}

	if 0 == CONFIG.MTU {
		CONFIG.MTU = 1472
		update = true
	}

	if 0 == CONFIG.PORT {
		CONFIG.PORT = uint16(rand.Int31n(20000)) + 10000
		update = true
	}

	port := CONFIG.PORT
	for ; port < 49151; port++ {
		if !net.PortUsed(port) {
			break
		}
	}

	if port != CONFIG.PORT {
		CONFIG.PORT = port
		update = true
	}

	if update {
		b := &bytes.Buffer{}
		Throw(toml.NewEncoder(b).Encode(CONFIG))
		Throw(ioutil.WriteFile(filepath, b.Bytes(), 0644))
	}
}
