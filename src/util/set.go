package util

type Set struct {
	tree *RBTree
}

func (set *Set) Insert(val int) {
	set.tree.Insert(val)
}

func (set *Set) Contains(val int) bool {
	if node := set.tree.Search(val); node != nil {
		return true
	}
	return false
}
