package main

import (
	"errors"
	"fmt"
	"os"
)

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func IsEmpty(s string) bool {
	return s == ""
}

func ClearConsole() {
	fmt.Print("\x0c") // Clear screen and print field.
}
