package util

import (
	"sync"
)

type Wait struct {
	sync.WaitGroup
}

func (w *Wait) Run(cb func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		cb()
	}()
}
