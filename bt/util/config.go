package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	. "github.com/urwork/throw"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ID  string
	MTU int
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
		CONFIG.ID = B64uuid()
		update = true
	}

	if 0 == CONFIG.MTU {
		CONFIG.MTU = 1472
		update = true
	}

	if update {
		b := &bytes.Buffer{}
		Throw(toml.NewEncoder(b).Encode(CONFIG))
		Throw(ioutil.WriteFile(filepath, b.Bytes(), 0644))
	}
}
