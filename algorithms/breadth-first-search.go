/*
In the Breadth-First Search (BFS) algorithm, we explore the nodes of a graph or a tree level by level, visiting all the neighboring nodes before moving to the next level. We use a queue data structure to keep track of the nodes to be visited.

In the example, we create a binary tree and perform a BFS traversal starting from the root node. The BFS function takes the root node as input and performs the BFS traversal. It uses a queue to enqueue the nodes and dequeue them one by one, printing their values. The traversal visits the nodes in the order 1, 2, 3, 4, 5, 6.

Output:
BFS Traversal:
1 2 3 4 5 6
*/

package main

import (
	"fmt"
)

type Node struct {
	Value    int
	Children []*Node
}

func BFS(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", current.Value)

		for _, child := range current.Children {
			queue = append(queue, child)
		}
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

	fmt.Println("BFS Traversal:")
	BFS(root)
}