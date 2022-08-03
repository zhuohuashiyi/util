package util

type LRU struct {
	Capacity int
	Number   int
	record   map[int]int
	list     List
}

func NewLRU(cap int) *LRU {
	return &LRU{Capacity: cap, Number: 0, record: make(map[int]int), list: List{Num: 0, Head: nil, Tail: nil}}
}

func (lru *LRU) Get(key int) int {
	num, ok := lru.record[key]
	if ok {
		lru.list.Delete(key)
		lru.list.Add(key)
		return num
	} else {
		return -1
	}
}

func (lru *LRU) Set(key, value int) {
	if lru.Number == lru.Capacity {
		delete(lru.record, lru.list.Head.Val)
		lru.list.Delete(lru.list.Head.Val)
		lru.Number--
	}
	_, ok := lru.record[key]
	if ok {
		lru.list.Delete(key)
		lru.list.Add(key)
		lru.record[key] = value
	} else {
		lru.list.Add(key)
		lru.record[key] = value
		lru.Number++
	}
}
