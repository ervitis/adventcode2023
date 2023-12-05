package common

import (
	"bufio"
	"os"
)

func GetScanner(fileName string) *bufio.Scanner {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(f)
}

func GetFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return f
}
