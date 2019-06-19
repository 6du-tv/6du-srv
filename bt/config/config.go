package config

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"path"

	udp "bt/udp"
	cryptoRand "crypto/rand"

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/sha3"

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

type Key struct {
	SECRET *ecdsa.PrivateKey
	HASH   [64]byte
}

var CONFIG Config
var KEY Key
var ROOT = Root()

func init() {
	initConfig()
	initKey()
}

func initKey() {
	filepath := path.Join(ROOT, CONFIG.PATH.KEY)
	secret, err := ioutil.ReadFile(filepath)

	curve := btcec.S256()
	if os.IsNotExist(err) || (err == nil && len(secret) == 0) {

		// 参考 https://github.com/ethereum/go-ethereum/blob/master/crypto/crypto.go
		// privateKey := hex.EncodeToString()
		// address := crypto.PubkeyToAddress(key.PublicKey).Hex()

		key, _ := ecdsa.GenerateKey(curve, cryptoRand.Reader)

		secret = key.D.Bytes()

		Throw(ioutil.WriteFile(filepath, secret, 0400))

	} else {
		if 32 != len(secret) {
			panic(errors.New("secret length != 32"))
		}
	}

	pk := secret
	x, y := curve.ScalarBaseMult(pk)

	pub := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	private := &ecdsa.PrivateKey{
		PublicKey: *pub,
		D:         new(big.Int).SetBytes(pk),
	}

	KEY.SECRET = private

	pubkey := btcec.PublicKey(*pub)
	h := make([]byte, 64)
	sha3.ShakeSum256(h, pubkey.SerializeCompressed())

	copy(KEY.HASH[:], h)
	fmt.Printf("%x\n", h)
	fmt.Printf("%x\n", KEY.HASH)

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
