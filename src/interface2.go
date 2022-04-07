package main

import (
	"errors"
	"strconv"
)

type Failure int

func (t Failure) Error() string {
	return strconv.Itoa(int(t))
}

func PrintError(err error) {
	println("error: " + err.Error())
}

func main() {
	PrintError(errors.New("This is a test!"))
	PrintError(Failure(42))
}
