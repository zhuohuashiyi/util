package util

type BinarySearchTree struct {
	Tree
}

func (tree *BinarySearchTree) Search(val int) *TreeNode {
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

func (tree *BinarySearchTree) Insert(val int) {
	traverse := tree.Head
	var father *TreeNode
	var left int
	for traverse != nil {
		father = traverse
		if traverse.Val > val {
			left = 1
			traverse = traverse.Left
		} else if traverse.Val < val {
			left = 0
			traverse = traverse.Right
		}
	}
	if father == nil {
		tree.Head = &TreeNode{Val: val}
	} else {
		traverse = &TreeNode{Val: val}
		if left == 0 {
			father.Right = traverse
		} else {
			father.Left = traverse
		}
	}
}

func (tree *BinarySearchTree) Delete(val int) {
	var father *TreeNode
	var left int
	traverse := tree.Head
	for traverse != nil {
		if traverse.Val > val {
			father = traverse
			left = 1
			traverse = traverse.Left
		} else if traverse.Val < val {
			father = traverse
			left = 0
			traverse = traverse.Right
		} else {
			if father != nil {
				del(father, traverse, left)
				return
			} else {
				tree.Head = remove(tree.Head)
				return
			}

		}

	}
}

func del(father *TreeNode, traverse *TreeNode, left int) {
	if traverse.Left == nil && traverse.Right == nil {
		if left == 1 {
			father.Left = nil
		} else {
			father.Right = nil
		}
	} else if traverse.Left == nil {
		if left == 1 {
			father.Left = remove(traverse)
		} else {
			father.Right = remove(traverse)
		}
	} else if traverse.Right == nil {
		if left == 1 {
			father.Left = remove(traverse)
		} else {
			father.Right = remove(traverse)
		}
	} else {
		if left == 1 {
			father.Left = remove(traverse)
		} else {
			father.Right = remove(traverse)
		}
	}
}

func remove(head *TreeNode) *TreeNode {
	var father, left, right *TreeNode
	if head.Left == nil && head.Right == nil {
		return nil
	} else if head.Right == nil {
		traverse := head.Left
		for traverse.Right != nil {
			father = traverse
			traverse = traverse.Right
		}
		left = head.Left
		right = head.Right
		head = traverse

		head.Right = right
		if father != nil {
			father.Left = remove(traverse)
			head.Left = left
		}
		return head
	} else {
		traverse := head.Right
		for traverse.Left != nil {
			father = traverse
			traverse = traverse.Left
		}
		left = head.Left
		right = head.Right
		head = traverse
		head.Left = left
		//head.Right = right
		if father != nil {
			father.Right = remove(traverse)
			head.Right = right
		}
		return head

	}

}
