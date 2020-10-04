package file

import (
	"fmt"
	"os"
	"reflect"
)

type File struct {
	data string
}

func NewFile() *File {
	return &File{}
}

func (f *File) Read() string {
	file, err := os.Open("Test.txt")
	if os.IsNotExist(err) {
		os.Create("Test.txt")
	}
	defer file.Close()

	//999 количество байт которое читаем из файла
	data := make([]byte, 999)
	n, err := file.Read(data)
	return string(data[:n])
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
