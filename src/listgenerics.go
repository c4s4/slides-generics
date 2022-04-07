package main

type Element[T any] struct {
	Next  *Element[T]
	Value T
}

type List[T any] struct {
	Front *Element[T]
	Last  *Element[T]
}

func (l *List[T]) PushBack(value T) {
	node := &Element[T]{
		Next:  nil,
		Value: value,
	}
	if l.Front == nil {
		l.Front = node
		l.Last = node
	} else {
		l.Last.Next = node
		l.Last = node
	}
}

func main() {
	list := &List[int]{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	sum := 0
	for n := list.Front; n != nil; n = n.Next {
		sum += n.Value
	}
	println(sum)
}
