package main

import "golang.org/x/exp/constraints"

func Max[N constraints.Ordered](x, y N) N {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(Max(1, 2))
	println(Max(1.2, 2.1))
	println(Max("abc", "def"))
}
