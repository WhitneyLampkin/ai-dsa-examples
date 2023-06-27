package main

import "fmt"

func main() {
    // Creating an array of integers
    arr := [5]int{1, 2, 3, 4, 5}

    // Accessing elements
    fmt.Println(arr[0]) // Output: 1

    // Updating an element
    arr[3] = 10
    fmt.Println(arr) // Output: [1 2 3 10 5]
}