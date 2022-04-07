package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(Max(1, 2))
}
