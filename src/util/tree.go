package util

type Tree struct {
	Head *TreeNode
	Size int
}

func PreOrder(head *TreeNode, arr *[]int) {
	if head == nil {
		return
	}
	*arr = append(*arr, head.Val)
	if head.Left != nil {
		PreOrder(head.Left, arr)
	}
	if head.Right != nil {
		PreOrder(head.Right, arr)
	}
}

func MidOrder(head *TreeNode, arr *[]int) {
	if head == nil {
		return
	}
	if head.Left != nil {
		MidOrder(head.Left, arr)
	}
	*arr = append(*arr, head.Val)
	if head.Right != nil {
		MidOrder(head.Right, arr)
	}
}

func PostOrder(head *TreeNode, arr *[]int) {
	if head == nil {
		return
	}
	if head.Left != nil {
		PostOrder(head.Left, arr)
	}
	if head.Right != nil {
		PostOrder(head.Right, arr)
	}
	*arr = append(*arr, head.Val)
}
