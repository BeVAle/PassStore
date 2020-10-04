package command

import "fmt"

type Update struct {
}

func (Update) Run(arguments []string) {
	fmt.Println("Update successfully!")
}