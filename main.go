package main

import (
	"github.com/bevale/passStore/src/controller"
	"github.com/bevale/passStore/src/file"
	"os"
)

func main()  {
	arguments := os.Args
	cliController := controller.NewCliController(arguments)
	cliController.Run()

	file1 := file.NewFile()
	file1.Read()






}
