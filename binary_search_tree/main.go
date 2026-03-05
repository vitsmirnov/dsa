package main

type BSTreeNode struct {
	val         int
	left, right *BSTreeNode
	// size int
}

type BSTree struct {
	root *BSTreeNode
	// size int
}

func MakeBSTree() BSTree {
	return BSTree{root: nil}
}

func (t *BSTree) Add(val int) {
	// time: O(h), space: O(1)

	node := t.findNode(val)
	if *node == nil {
		*node = &BSTreeNode{val: val, left: nil, right: nil}
	}
}

func (t *BSTree) Remove(val int) {
	// time: O(h), space: O(1)

	node := t.findNode(val)
	if *node == nil {
		return
	}

	if (**node).right == nil {
		*node = (**node).left
	} else if (**node).left == nil {
		*node = (**node).right
	} else {
		minRight := &(**node).right
		for (**minRight).left != nil {
			minRight = &(**minRight).left
		}
		(**node).val = (**minRight).val
		*minRight = (**minRight).right
	}
}

func (t *BSTree) Exist(val int) bool {
	// time: O(h), space: O(1)

	return *t.findNode(val) != nil
}

func (t *BSTree) findNode(val int) **BSTreeNode {
	// time: O(h), space: O(1)

	node := &t.root
	for *node != nil && (*node).val != val {
		if val < (*node).val {
			node = &(*node).left
		} else { // val > (*node).val
			node = &(*node).right
		}
	}
	return node
}

func (t *BSTree) Add1(val int) {
	if t.root == nil {
		t.root = &BSTreeNode{val: val, left: nil, right: nil}
		return
	}
	node := t.root
	for {
		if val < node.val {
			if node.left == nil {
				node.left = &BSTreeNode{val: val, left: nil, right: nil}
				return
			} else {
				node = node.left
			}
		} else if val > node.val {
			if node.right == nil {
				node.right = &BSTreeNode{val: val, left: nil, right: nil}
				return
			} else {
				node = node.right
			}
		} else {
			return
		}
	}
}

func (t *BSTree) Exist1(val int) bool {
	node := t.root
	for node != nil {
		if val < node.val {
			node = node.left
		} else if val > node.val {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

// temp

func deleteNode(root *TreeNode, key int) *TreeNode {
	// time: O(h), space: O(1)

	node := findNode(&root, key)
	if *node == nil {
		return root
	}

	if (**node).Right == nil {
		*node = (**node).Left
	} else if (**node).Left == nil {
		*node = (**node).Right
	} else {
		minRight := &(**node).Right
		for (**minRight).Left != nil {
			minRight = &(**minRight).Left
		}
		(**node).Val = (**minRight).Val
		*minRight = (**minRight).Right
	}
	return root
}

func findNode(root **TreeNode, key int) **TreeNode {
	// time: O(h), space: O(1)

	for *root != nil && (**root).Val != key {
		if key < (**root).Val {
			root = &(**root).Left
		} else { // key > (**root).Val
			root = &(**root).Right
		}
	}
	return root
}

//  temp

// func addBSTNode(root *BSTreeNode, val int) *BSTreeNode {
//     if root == nil {
//         return &BSTreeNode{
//             val: val,
//             left: nil,
//             right: nil}
//     }
//     if val < root.val {
//         return addBSTNode(root.left, val)
//     } else if val > root.val {
//         return addBSTNode(root.right, val)
//     } else {
//         return root
//     }
// }
