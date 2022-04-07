package main

import "golang.org/x/exp/constraints"

func Max[N constraints.Ordered](x, y N) N {
	if x > y { return x }
	return y
}

func main() {
	println(Max("abc", "def"))
}
