package util

import (
	"strconv"
	"strings"
)

type List struct {
	Num  int
	Head *ListNode
	Tail *ListNode
}

func (list *List) Delete(val int) {
	var traverse *ListNode = list.Head
	for traverse != nil && traverse.Val != val {
		traverse = traverse.Next
	}
	if traverse != nil {
		list.Num--
		if traverse == list.Head {
			list.Head = traverse.Next
		}
		if traverse == list.Tail {
			list.Tail = traverse.Before
		}
		if traverse.Before != nil {
			traverse.Before.Next = traverse.Next
		}
		if traverse.Next != nil {
			traverse.Next.Before = traverse.Before
		}
	}
}

func (list *List) Add(val int) *ListNode {
	list.Num++
	temp := ListNode{Val: val, Next: nil, Before: nil}
	if list.Head == nil {
		list.Head = &temp
		list.Tail = &temp
	} else {
		list.Tail.Next = &temp
		temp.Before = list.Tail
		list.Tail = &temp
	}
	return &temp
}

func (list *List) String() string {
	res := make([]string, 0)
	for traverse := list.Head; traverse != nil; traverse = traverse.Next {
		res = append(res, strconv.Itoa(traverse.Val))
	}
	ans := strings.Join(res, " ")
	return ans

}
