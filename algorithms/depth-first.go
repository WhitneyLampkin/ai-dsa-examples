/*
In the Depth-First Search (DFS) algorithm, we explore the nodes of a graph or a tree by traversing as deep as possible before backtracking. We recursively visit all the adjacent nodes before moving to the next adjacent node.

In the example, we create a binary tree and perform a DFS traversal starting from the root node. The DFS function takes the root node as input and performs the DFS traversal. It recursively visits each node, printing its value. The traversal visits the nodes in the order 1, 2, 4, 5, 3, 6.

Output:
	DFS Traversal:
	1 2 4 5 3 6
*/

package main

import "fmt"

type Node struct {
	Value    int
	Children []*Node
}

func DFS(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Value)

	for _, child := range root.Children {
		DFS(child)
	}
}

func main() {
	// Create a binary tree
	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	node6 := &Node{Value: 6}

	root.Children = []*Node{node2, node3}
	node2.Children = []*Node{node4, node5}
	node3.Children = []*Node{node6}

	fmt.Println("DFS Traversal:")
	DFS(root)
}