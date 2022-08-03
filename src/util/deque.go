package util

type Deque struct {
	Element []int
	Len int
}

func (deque *Deque) Add(val int) {
	deque.Element = append(deque.Element, val)
	deque.Len++
}

func (deque *Deque) Remove() int {
	res := deque.Element[0]
	deque.Element = deque.Element[1: deque.Len]
	deque.Len--
	return res
}

func (deque *Deque) Size() int {
	return deque.Len
}

func (deque *Deque) IsEmpty() bool {
	return deque.Len == 0
}
