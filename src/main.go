package main

import (
	"util/src/util"
)

func main() {
	var tree *util.RBTree = &util.RBTree{}
	tree.Insert(11)
	tree.Insert(9)
	tree.Insert(7)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(0)
	tree.Delete(1)

}
