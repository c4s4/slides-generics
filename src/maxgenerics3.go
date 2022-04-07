package main

type Number interface {
	~int | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Truc int

func Max[N Number](x, y N) N {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(Max(Truc(1), Truc(2)))
}
