package command

import (
	"fmt"
	"github.com/bevale/passStore/src/crypto"
	"github.com/bevale/passStore/src/file"
	"log"
)

type Search struct {
}

func (Search) Run(arguments []string) {
	file1 := file.NewFile()

	encodeString := file1.Read()
	key := []byte("1234567890abcdef") // 32 bytes
	result, err := crypto.Decrypt(key, []byte(encodeString))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", result)

}
