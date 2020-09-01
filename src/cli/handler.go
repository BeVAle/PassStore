package cli

import "fmt"

type ArgsHandler struct {
	arguments []string
	keys      []uint8
	values    []string
}

func NewArgsHandler(arguments []string) *ArgsHandler {
	return &ArgsHandler{arguments: arguments}
}

func (arg *ArgsHandler) SetKeyFromCli() {
	for _, value := range arg.arguments {
		if value[0] == '-' {
			for x := 1; x < len(value); x++ {
				arg.keys = append(arg.keys, value[x])
			}
		}
	}
	arg.showKeys()
}

func (arg *ArgsHandler) showKeys() {
	for _, v := range arg.keys {
		fmt.Printf("%c\n", v)
	}
}

func (arg *ArgsHandler) DoHandler() {
	if len(arg.arguments) == 1 {
		fmt.Println("Welcome! Type some key for work (＾▽＾). (-h for Help ヽ(・∀・)ﾉ )")
		return
	}
	arg.SetKeyFromCli()
}
