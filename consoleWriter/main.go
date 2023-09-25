package main

import (
	"fmt"
	"io"
	"os"
)

type consoleWriter struct{}

func (cw consoleWriter) Write(buf []byte) (n int, err error) {
	fmt.Println(string(buf))

	return len(buf), nil
}

func main() {
	fileName := os.Args[1]
	fmt.Println("Reading from file", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file", fileName, err)
	}

	io.Copy(consoleWriter{}, file)

}
