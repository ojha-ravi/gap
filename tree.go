package main

import (
	"fmt"
)

type node struct {
	lhs *node
	rhs *node
	val int
}

func createNode(val int) *node {
	newNode := &node{}
	newNode.val = val
	return newNode
}

func createBST(a []int) *node {
	var root *node

	for _, v := range a {
		inserNodeIntoBST(&root, v)
	}
	return root
}

func inserNodeIntoBST(root **node, val int) {
	if *root == nil {
		*root = createNode(val)
	} else if val > (*root).val {
		if (*root).rhs == nil {
			(*root).rhs = createNode(val)
		} else {
			inserNodeIntoBST(&(*root).rhs, val)
		}
	} else {
		if (*root).lhs == nil {
			(*root).lhs = createNode(val)
		} else {
			inserNodeIntoBST(&(*root).lhs, val)
		}
	}
}

func inOrder(root *node, r *[]int) {
	if root == nil {
		return
	}

	inOrder(root.lhs, r)
	*r = append(*r, root.val)
	inOrder(root.rhs, r)
}

func preOrder(root *node, r *[]int) {
	if root == nil {
		return
	}

	*r = append(*r, root.val)
	preOrder(root.lhs, r)
	preOrder(root.rhs, r)
}

func postOrder(root *node, r *[]int) {
	if root == nil {
		return
	}

	postOrder(root.lhs, r)
	postOrder(root.rhs, r)
	*r = append(*r, root.val)
}

// func inOrderWithoutRec(root *node)  {
// 	for root != nil {
// 		fm
// 	}
// }

func mirror(root **node) {
	if *root == nil {
		return
	}
	mirror(&((*root).lhs))
	mirror(&((*root).rhs))
	(*root).lhs, (*root).rhs = (*root).rhs, (*root).lhs
}

func main() {
	var res []int
	tree := createBST([]int{1, 2, 3, 4, 5, 6})
	inOrder(tree, &res)
	fmt.Println(res)

	res = []int{}
	mirror(&tree)
	inOrder(tree, &res)
	fmt.Println(res)
}
