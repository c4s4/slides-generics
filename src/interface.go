package main

import "fmt"

func Repeat(something interface{}, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(something)
	}
}

func main() {
	Repeat("Hello World!", 3)
    Repeat(42, 3)
}
