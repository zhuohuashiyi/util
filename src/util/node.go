package util

type ListNode struct {
	Val    int
	Next   *ListNode
	Before *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type RBTreeNode struct {
	Val    int
	Left   *RBTreeNode
	Right  *RBTreeNode
	Father *RBTreeNode
	red    int
	left   int
}
