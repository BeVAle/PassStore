package file

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type File struct {
	data string
}

func NewFile() *File {
	return &File{}
}

func (f *File) Read() {
	file, err := os.Open("Test.txt")
	if os.IsNotExist(err){
		os.Create("Test.txt")
	}
	defer file.Close()

	data := make([]byte, 64)

	for  {
		n, err := file.Read(data)
		if err == io.EOF{
			break
		}
		fmt.Print(string(data[:n]))
	}
}

func (f *File) Write(text string) {
	file, err := os.Create("Test.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(reflect.TypeOf(file).String())
	file.WriteString(text)
}

