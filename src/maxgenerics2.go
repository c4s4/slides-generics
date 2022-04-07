package main

type Number interface {
	int | int16 | int32 | int64 | float32 | float64
}

func Max[N Number](x, y N) N {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(Max(1, 2))
	println(Max(1.2, 2.1))
}
