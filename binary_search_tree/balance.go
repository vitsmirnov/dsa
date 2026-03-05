package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// varsion 2 (Day-Stout-Warren algorithm (in place))

func balanceBST2(root *TreeNode) *TreeNode {
	// time: O(n), space: O(1)

	makeRightVine3(&root)
	nodeCount := countVineNodes(root)
	m := 1
	for m < nodeCount+1 {
		m <<= 1
	}
	m = (m >> 1) - 1 // number of nodes in the closest perfectly balanced tree

	doLeftRotations3(&root, nodeCount-m)
	for ; m > 1; m /= 2 {
		doLeftRotations3(&root, m/2)
	}
	return root
}

func doLeftRotations3(root **TreeNode, count int) {
	for range count {
		*root = rotateLeft(*root)
		root = &(*root).Right
	}
}

func doLeftRotations2(root *TreeNode, count int) *TreeNode {
	dummy := &TreeNode{Left: nil, Right: root}
	cur := dummy
	for range count {
		cur.Right = rotateLeft(cur.Right)
		cur = cur.Right
	}
	return dummy.Right
}

func doLeftRotations1(root *TreeNode, count int) *TreeNode {
	root = rotateLeft(root)
	cur := root
	for range count - 1 {
		cur.Right = rotateLeft(cur.Right)
		cur = cur.Right
	}
	return root
}

func makeRightVine3(root **TreeNode) {
	for ; *root != nil; root = &(*root).Right {
		for (*root).Left != nil {
			*root = rotateRight(*root)
		}
	}
}

func makeRightVine2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{Left: nil, Right: root}
	for node := dummy; node.Right != nil; {
		for node.Right.Left != nil {
			node.Right = rotateRight(node.Right)
		}
		node = node.Right
	}
	return dummy.Right
}

func makeRightVine1(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = rotateRight(root)
	}
	for cur := root; cur.Right != nil; cur = cur.Right {
		for cur.Right.Left != nil {
			cur.Right = rotateRight(cur.Right)
		}
	}
	return root
}

func rotateRight(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	res := root.Left
	right := res.Right
	res.Right = root
	root.Left = right
	return res
}

func rotateLeft(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	res := root.Right
	left := res.Left
	res.Left = root
	root.Right = left
	return res
}

func countVineNodes(root *TreeNode) int {
	count := 0
	for ; root != nil; root = root.Right {
		count++
	}
	return count
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// version 1 (reconstruction: inorder traversal + recursive construction)

func balanceBST(root *TreeNode) *TreeNode {
	// time: O(n), space: O(n)

	return balance(treeToArray(root))
}

func balance(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	return &TreeNode{
		Val:   nums[i],
		Left:  balance(nums[:i]),
		Right: balance(nums[i+1:])}
}

func treeToArray(root *TreeNode) []int {
	res := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			res = append(res, root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}
