package keys

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/chuckleheads/habitat/components/hab/pkg/crypto"
)

type KeyType int

const (
	Sig KeyType = iota
	Box
	Sym
)

type Key struct {
	Name   string
	Rev    string
	Type   KeyType
	Key    []byte
	Secret bool
}

type KeyPair struct {
	PublicKey  *Key
	PrivateKey *Key
}

func (k *Key) nameWithRev() string {
	return fmt.Sprintf("%s-%s", k.Name, k.Rev)
}

func KeyFromString(keystr string) Key {
	scanner := bufio.NewScanner(strings.NewReader(keystr))
	key := Key{}
	// Key Type
	if !scanner.Scan() {
		panic(fmt.Sprintf("Malformed key string:\n(%s)", keystr))
	}

	if scanner.Text() == crypto.PublicSigKeyVersion || scanner.Text() == crypto.PublicBoxKeyVersion {
		key.Secret = false
	} else if scanner.Text() == crypto.PrivateSigKeyVersion ||
		scanner.Text() == crypto.PrivateSymKeyVersion ||
		scanner.Text() == crypto.PrivateSigKeyVersion ||
		scanner.Text() == crypto.PrivateBoxKeyVersion {
		key.Secret = true
	} else {
		panic(fmt.Sprintf("Unsupported Key Version: %s", scanner.Text()))
	}

	// Name and Rev
	if scanner.Scan() {
		nameAndRev := strings.Split(scanner.Text(), "-")
		key.Name, key.Rev = nameAndRev[0], nameAndRev[1]
	} else {
		panic(fmt.Sprintf("Malformed key string:\n(%s)", keystr))
	}

	// Key Content
	scanner.Scan()
	if !scanner.Scan() {
		panic(fmt.Sprintf("Malformed key string:\n(%s)", keystr))
	}

	keyBytes, err := base64.StdEncoding.DecodeString(scanner.Text())
	if err != nil {
		panic(err.Error())
	}

	key.Key = keyBytes
	return key
}
