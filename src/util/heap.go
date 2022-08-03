package util

import (
	"strconv"
)

type Heap struct {
	Element []int
	Len    int
}

func (heap *Heap) Add(val int) {
	heap.Element = append(heap.Element, val)
	heap.siftUp(heap.Len)
	heap.Len++

}

func (heap *Heap) Remove() int {
	res := heap.Element[heap.Len-1]
	heap.Len--
	heap.Element = heap.Element[0:heap.Len]
	return res
}

func (heap *Heap) Set(idx int, val int) {
	heap.Element[idx] = val
}
func (heap *Heap) Top() int {
	return heap.Element[0]
}

func (heap *Heap) Size() int {
	return heap.Len
}
func (heap *Heap) siftUp(pos int) {
	i := pos
	j := (pos - 1) / 2
	for j >= 0 {
		if heap.Element[j] > heap.Element[i] {
			heap.Element[j], heap.Element[i] = heap.Element[i], heap.Element[j]
			i = j
			j = (j - 1) / 2
		} else {
			break
		}
	}
}

func (heap *Heap) siftDown() {
	i := 0
	j1 := 2*i + 1
	j2 := 2*i + 2
	for j1 < heap.Len {
		if (j2 >= heap.Len || heap.Element[j1] <= heap.Element[j2]) && heap.Element[j1] < heap.Element[i] {
			heap.Element[j1], heap.Element[i] = heap.Element[i], heap.Element[j1]
			i = j1
		} else if j2 < heap.Len && heap.Element[j2] < heap.Element[j1] && heap.Element[j2] < heap.Element[i] {
			heap.Element[j2], heap.Element[i] = heap.Element[i], heap.Element[j2]
			i = j2
		} else {
			break
		}
		j1 = 2*i + 1
		j2 = 2*i + 2
	}

}

func (heap *Heap) String() string {
	var res string
	res += "["
	for _, v := range heap.Element {
		res += strconv.Itoa(v)
		res += ", "
	}
	res = res[0 : len(res)-2]
	res += "]"
	return res
}

