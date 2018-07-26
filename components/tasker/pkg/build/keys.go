package build

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/chuckleheads/habitat/components/builder-depot-client/client"
	"github.com/chuckleheads/habitat/components/hab/pkg/crypto/keys"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/access_token"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

func fetchOriginKey(origin string) ([]byte, string) {
	tokenLife := time.Duration(2) * time.Hour
	client := client.New("https://bldr.acceptance.habitat.sh", generateAccessToken(0, uint32(7), tokenLife))
	sk, name, err := client.FetchOriginSecretKey(origin)
	if err != nil {
		log.Fatalln(err)
	}
	return sk, name
}

func writeKeyToDisk(key []byte, filename string) {
	err := ioutil.WriteFile(filepath.Join("/", "hab", "cache", "keys", filename), key, 0600)
	if err != nil {
		log.Fatalln(err)
	}
}

func generateAccessToken(accountID uint64, flags uint32, lifetime time.Duration) string {
	expires := time.Now().Add(lifetime).Unix()

	var token = access_token.AccessToken{
		AccountId: accountID,
		Flags:     flags,
		Expires:   expires,
	}

	tokenBytes, err := proto.Marshal(&token)
	if err != nil {
		log.Fatalln("Failed to encode Token:", err)
	}

	kp := keys.NewFromString(viper.GetString("bldr.public"), viper.GetString("bldr.private"))

	secret, err := kp.Encrypt(tokenBytes, &kp)
	if err != nil {
		log.Fatalln("Failed to encode AccessToken: ", err)
	}

	ciphertext := base64.StdEncoding.EncodeToString([]byte(secret))

	return fmt.Sprintf("%s%s", AccessTokenPrefix, ciphertext)
}
