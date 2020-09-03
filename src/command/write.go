package command

import (
	"github.com/bevale/passStore/src/crypto"
	"github.com/bevale/passStore/src/file"
	"strings"
)

type Write struct {
}

func (Write) Run(arguments []string) {
	record := strings.Join(arguments[:], ", ")
	key := []byte("1234567890abcdef") // 32 bytes
	plaintext := []byte(record)
	text, err :=  crypto.Encrypt(key, plaintext)
	if err != nil {
		panic(err)
	}

	(file.NewFile()).Write(string(text))




}
