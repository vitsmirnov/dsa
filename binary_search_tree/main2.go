package main

import "fmt"

type BSTreeNode2 struct {
	val         int
	left, right *BSTreeNode2
	count       int // val count
	size        int // subtree size
}

func MakeBSTNode2(val int) *BSTreeNode2 {
	return &BSTreeNode2{
		val:   val,
		left:  nil,
		right: nil,
		count: 1,
		size:  1}
}

func (n *BSTreeNode2) Size() int {
	if n == nil {
		return 0
	}
	return n.size
}

func (n *BSTreeNode2) Count() int {
	if n == nil {
		return 0
	}
	return n.size
}

type BSTree2 struct {
	root *BSTreeNode2
	// size int
}

func MakeBSTree2() BSTree2 {
	return BSTree2{root: nil}
}

func (t *BSTree2) Add(val int) {
	if t.root == nil {
		t.root = MakeBSTNode2(val)
		return
	}
	node := t.root
	for {
		node.size++
		if val < node.val {
			if node.left == nil {
				node.left = MakeBSTNode2(val)
				return
			} else {
				node = node.left
			}
		} else if val > node.val {
			if node.right == nil {
				node.right = MakeBSTNode2(val)
				return
			} else {
				node = node.right
			}
		} else {
			node.count++
			return
		}
	}
}

func (t *BSTree2) Remove(val int) {
	// time: O(h), space: O(1)

	if !t.Exist(val) {
		return
	}
	node := &t.root
	for (*node).val != val {
		(*node).size--
		if val < (*node).val {
			node = &(*node).left
		} else {
			node = &(*node).right
		}
	}
	(*node).size--
	(*node).count--
	if (*node).count > 0 {
		return
	}

	if (*node).right == nil {
		*node = (*node).left
	} else if (*node).left == nil {
		*node = (*node).right
	} else {
		minRight := &(*node).right
		for (*minRight).left != nil {
			minRight = &(*minRight).left
		}
		count := (*minRight).count
		(*node).val = (*minRight).val
		(*node).count = (*minRight).count
		minRight = &(*node).right
		for (*minRight).left != nil {
			(*minRight).size -= count
			minRight = &(*minRight).left
		}
		*minRight = (*minRight).right
	}
}

func (t *BSTree2) CountVal(val int) int {
	return (*t.findNode(val)).Count()
}

func (t *BSTree2) CountLessThan(val int) int {
	count := 0
	node := t.root
	for node != nil {
		if val < node.val {
			node = node.left
		} else if val > node.val {
			// count += node.size - node.right.Size()
			count += node.count + node.left.Size()
			node = node.right
		} else {
			count += node.left.Size()
			break
		}
	}
	return count
}

func (t *BSTree2) CountGreaterThan(val int) int {
	return t.root.Size() - t.CountVal(val) - t.CountLessThan(val)
}

func (t *BSTree2) Exist(val int) bool {
	// time: O(h), space: O(1)

	return *t.findNode(val) != nil
}

func (t *BSTree2) findNode(val int) **BSTreeNode2 {
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

///

func (t *BSTree2) _Add1(val int) {
	// time: O(h), space: O(1)

	node := t.findNode(val)
	if *node == nil {
		*node = &BSTreeNode2{val: val, left: nil, right: nil}
	}
}

func (t *BSTree2) _Remove1(val int) {
	// time: O(h), space: O(1)

	node := t.findNode(val)
	if *node == nil {
		return
	}

	// node.count

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

func main() {
	n := &BSTreeNode2{
		val:   7,
		count: 2,
		size:  4,
	}
	n = nil
	fmt.Println(n.Size())
}
