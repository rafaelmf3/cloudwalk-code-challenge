package internal

import (
	"os"
)

func ReadFile(fileNamePath string) *os.File {
	f, err := os.Open(fileNamePath)
	if err != nil {
		panic(err)
	}
	return f
}

func CreateFile(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}
