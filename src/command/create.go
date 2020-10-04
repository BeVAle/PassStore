package command

import (
	"encoding/json"
	"fmt"
	"github.com/bevale/passStore/src/crypto"
	"github.com/bevale/passStore/src/file"
	"github.com/bevale/passStore/src/model"
	"log"
)

type Create struct {
}

func (Create) Run(arguments []string) {

	passRecord1 := model.PassRecord{
		Email: "bedrin-vlad@mail.ru",
		Nickname: "BeVAle",
		Website: "vk.com",
		Password: "123456",
		OtherData: "chenit' zakinu",
	}
	passRecord2 := model.PassRecord{
		Email: "bedrin-vlad@mail.ru",
		Nickname: "Lols",
		Website: "linkedin",
		Password: "12345678",
		OtherData: "chenit' zakinu",
	}

	passRecords := model.PassRecords{
		PassRecords: []model.PassRecord{passRecord1, passRecord2},
	}

	fmt.Println("Create successfully!")
	var jsonData []byte
	jsonData, err := json.Marshal(passRecords)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	key := []byte("1234567890abcdef") // 32 bytes
	plaintext := []byte(jsonData)
	text, err :=  crypto.Encrypt(key, plaintext)
	if err != nil {
		panic(err)
	}

	(file.NewFile()).Write(string(text))

}
