package main

import (
	"github.com/bevale/passStore/src/cli"
	"os"
)

func main()  {
	arguments := os.Args
	argsHandler :=  cli.NewArgsHandler(arguments)
	argsHandler.DoHandler()
}
