package util

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func B64uuid() string {
	token := make([]byte, 16)
	rand.Read(token)
	return base64.RawURLEncoding.EncodeToString(token)
}
