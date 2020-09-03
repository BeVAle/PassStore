package controller

import (
	"fmt"
	"github.com/bevale/passStore/src/command"
)

type CliController struct {
	arguments []string
}

func NewCliController(arguments []string) *CliController {
	return &CliController{arguments: arguments}
}

func (c *CliController) Run() {
	c.getCommand()
}

func (c *CliController) getCommand() {
	if len(c.arguments) == 1 {
		fmt.Println("Welcome! Type some key for work (＾▽＾). (-h for Help ヽ(・∀・)ﾉ )")
		return
	}
	switch c.arguments[1] {
	case "search":
		command.Command = command.Search{}
		break
	case "write":
		command.Command = command.Write{}
		break
	}

	args := make([]string, 2)
	copy(args, c.arguments[2:])
	command.Command.Run(args)

}
