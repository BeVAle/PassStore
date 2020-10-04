package command

import "fmt"

type Remove struct {
}

func (Remove) Run(arguments []string) {
	fmt.Println("Remove successfully!")
}
