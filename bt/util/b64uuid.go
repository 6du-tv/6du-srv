package util

import (
	"encoding/base64"
	"math/rand"
)

func B64uuid() string {
	token := make([]byte, 16)
	rand.Read(token)
	return base64.RawURLEncoding.EncodeToString(token)
}
