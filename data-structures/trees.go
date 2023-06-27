package main

import "fmt"

type Node struct {
    data  int
    left  *Node
    right *Node
}

func main() {
    // Creating a binary tree
    root := &Node{data: 1}
    root.left = &Node{data: 2}
    root.right = &Node{data: 3}

    // Accessing elements
    fmt.Println(root.data)        // Output: 1
    fmt.Println(root.left.data)   // Output: 2
    fmt.Println(root.right.data)  // Output: 3

    // Updating an element
    root.data = 5
    fmt.Println(root.data)        // Output: 5
}