package main

func Max[N int | float64](x, y N) N {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(Max(1, 2))
	println(Max(1.2, 2.1))
}
