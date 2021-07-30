package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return nil
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return ans
}

//106. 从中序与后序遍历序列构造二叉树
//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
func buildTree(inorder []int, postorder []int) *TreeNode {
	inOrderMap := make(map[int]int)

	for idx, val := range inorder {
		inOrderMap[val] = idx
	}

	var build func(int, int) *TreeNode
	build = func(inorderLeft int, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{
			Val: val,
		}
		inorderIdx := inOrderMap[val]
		root.Right = build(inorderIdx+1, inorderRight)
		root.Left = build(inorderLeft, inorderIdx-1)
		return root
	}
	return build(0, len(inorder)-1)
}

//105. 从前序与中序遍历序列构造二叉树
func buildTreeByPreorder(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for idx, val := range inorder {
		inorderMap[val] = idx
	}

	var build func(int, int) *TreeNode
	build = func(inorderLeft int, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		val := preorder[0]
		preorder = preorder[1:]
		root := &TreeNode{
			Val: val,
		}
		inorderIdx := inorderMap[val]
		root.Left = build(inorderLeft, inorderIdx-1)
		root.Right = build(inorderIdx+1, inorderRight)
		return root
	}
	return build(0, len(inorder)-1)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}
