package command

import "fmt"

type View struct {
}

func (View) Run(arguments []string) {
	fmt.Println("View successfully!")
}
