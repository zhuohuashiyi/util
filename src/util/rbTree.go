package util

type RBTree struct {
	Head *RBTreeNode
	size int
}

func (tree *RBTree) Search(val int) *RBTreeNode {
	traverse := tree.Head
	for traverse != nil {
		if traverse.Val > val {
			traverse = traverse.Left
		} else if traverse.Val < val {
			traverse = traverse.Right
		} else {
			return traverse
		}
	}
	return nil
}

func (tree *RBTree) Insert(val int) {
	var father, uncle *RBTreeNode
	var left int
	traverse := tree.Head
	for traverse != nil {
		father = traverse
		if traverse.Val > val {
			left = 1
			traverse = traverse.Left
		} else if traverse.Val < val {
			left = 0
			traverse = traverse.Right
		} else {
			return
		}
	}
	tree.size++
	if father == nil {
		tree.Head = &RBTreeNode{Val: val, red: 0}
	} else {
		traverse = &RBTreeNode{Val: val, red: 1, Father: father, left: left}
		if father.left == 1 {
			uncle = father.Father.Right
		} else {
			if father.Father == nil {
				uncle = nil
			} else {
				uncle = father.Father.Left
			}

		}
		if traverse.left == 0 {
			father.Right = traverse
		} else {
			father.Left = traverse
		}
		tree.siftUp(traverse, uncle, father)

	}
}

func (tree *RBTree) Delete(val int) {
	node := tree.Search(val)
	tree.del(node)
}

func (tree *RBTree) Size() int {
	return tree.size
}

func RBMidOrder(head *RBTreeNode, arr *[]int) {
	if head == nil {
		return
	}
	if head.Left != nil {
		RBMidOrder(head.Left, arr)
	}
	*arr = append(*arr, head.Val)
	if head.Right != nil {
		RBMidOrder(head.Right, arr)
	}
}

func (tree *RBTree) del(node *RBTreeNode) {
	tree.size--
	if node.Father == nil {
		tree.Head = nil
	} else if node.Left == nil && node.Right == nil {
		if node.left == 1 {
			node.Father.Left = nil
		} else {
			node.Father.Right = nil
		}
	} else if node.Left == nil || node.Right == nil {
		if node.left == 1 {
			if node.Left != nil {
				node.Father.Left = node.Left
				node.Left.Father = node.Father
				node.Left.left = node.left
				tree.siftDown(node.Left)
			} else {
				node.Father.Left = node.Right
				node.Right.Father = node.Father
				node.Right.left = node.left
				tree.siftDown(node.Right)
			}
		} else {
			if node.Left != nil {
				node.Father.Right = node.Left
				node.Left.Father = node.Father
				node.Left.left = node.left
				tree.siftDown(node.Left)
			} else {
				node.Father.Right = node.Right
				node.Right.Father = node.Father
				node.Right.left = node.left
				tree.siftDown(node.Right)
			}
		}
	} else {
		traverse := node.Right
		for traverse.Left != nil {
			traverse = traverse.Left
		}
		tree.del(traverse)
		traverse.Left = node.Left
		traverse.Right = node.Right
		tree.siftDown(traverse)
	}
}

func (tree *RBTree) siftUp(self, uncle, father *RBTreeNode) {
	if father == nil {
		self.red = 0
		return
	} else if father.red == 0 {
		return
	} else if uncle != nil && uncle.red == 1 {
		uncle.red = 0
		father.red = 0
		father.Father.red = 1
		self = father.Father
		father = self.Father
		if father != nil && father.Father != nil {
			if father.left == 1 {
				uncle = father.Father.Right
			} else {
				uncle = father.Father.Left
			}

		} else {
			uncle = nil
		}
		tree.siftUp(self, uncle, father)
	} else if father.left == 1 {
		if self.left == 1 {
			father.red = 0
			father.Father.red = 1
			tree.rightRotate(father.Father)
		} else {
			tree.leftRotate(father)
			father.Father.red = 0
			father.Father.Father.red = 1
			tree.rightRotate(father.Father.Father)
		}
	} else {
		if self.left == 1 {
			father.red = 0
			father.Father.red = 1
			tree.leftRotate(father.Father)
		} else {
			tree.rightRotate(father)
			father.Father.red = 0
			father.Father.Father.red = 1
			tree.rightRotate(father.Father.Father)
		}
	}

}

func (tree *RBTree) siftDown(self *RBTreeNode) {
	if self.red == 1 {
		self.red = 0
	} else {
		var brother *RBTreeNode
		if self.left == 1 {
			brother = self.Father.Right
			if brother.red == 1 {
				brother.red = 0
				brother.Father.red = 1
				tree.leftRotate(brother.Father)
				brother.Father.red = 1
				tree.siftDown(brother.Father.Father)
			} else {
				if brother.Right.red == 1 {
					brother.red = brother.Father.red
					brother.Father.red = 0
					brother.Right.red = 0
					tree.leftRotate(brother.Father)
				} else if brother.Right.red == 0 && brother.Left.red == 1 {
					brother.red = 1
					brother.Left.red = 0
					tree.rightRotate(brother)
					brother.Father.red = brother.Father.Father.red
					brother.Father.Father.red = 0
					brother.red = 0
					tree.leftRotate(brother.Father.Father)
				} else {
					brother.red = 1
					tree.siftDown(brother.Father)
				}
			}
		} else {
			brother = self.Father.Left
			if brother.red == 1 {
				brother.red = 0
				brother.Father.red = 1
				tree.rightRotate(brother.Father)
				brother.Father.red = 1
				tree.siftDown(brother.Father.Father)
			} else {
				if brother.Left.red == 1 {
					brother.red = brother.Father.red
					brother.Father.red = 0
					brother.Left.red = 0
					tree.rightRotate(brother.Father)
				} else if brother.Left.red == 0 && brother.Right.red == 1 {
					brother.red = 1
					brother.Right.red = 0
					tree.leftRotate(brother)
					brother.Father.red = brother.Father.Father.red
					brother.Father.Father.red = 0
					brother.red = 0
					tree.leftRotate(brother.Father.Father)
				} else {
					brother.red = 1
					tree.siftDown(brother.Father)
				}
			}
		}
	}
}

func (tree *RBTree) leftRotate(node *RBTreeNode) {
	temp := node.Right
	temp.Father = node.Father
	node.Right = temp.Left
	temp.Left = node
	temp.left = node.left
	if node.Father != nil {
		if node.left == 1 {
			node.Father.Left = temp
		} else {
			node.Father.Right = temp
		}
	} else {
		tree.Head = temp
	}
	node.Father = temp
	node.left = 1
	return
}

func (tree *RBTree) rightRotate(node *RBTreeNode) {
	temp := node.Left
	temp.Father = node.Father
	node.Left = temp.Right
	temp.Right = node
	temp.left = node.left
	if node.Father != nil {
		if node.left == 1 {
			node.Father.Left = temp
		} else {
			node.Father.Right = temp
		}
	} else {
		tree.Head = temp
	}
	node.Father = temp
	node.left = 0
	return
}
