package config

import (
	"os"
	"path"
	"runtime"

	. "github.com/urwork/throw"
)

func Root() string {
	root := os.Getenv("_6DU_ROOT")
	if len(root) == 0 {
		root = os.Getenv("HOME")

		if len(root) == 0 {
			_, filename, _, _ := runtime.Caller(0)
			root = path.Dir(filename)
		} else {
			root = path.Join(root, ".config", "6du")
		}
	}
	err := os.MkdirAll(root, os.ModePerm)
	Throw(err)
	return root
}
