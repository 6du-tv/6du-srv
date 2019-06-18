package util

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandByteB64(size int) string {
	token := make([]byte, size)
	rand.Read(token)
	return base64.RawURLEncoding.EncodeToString(token)
}
