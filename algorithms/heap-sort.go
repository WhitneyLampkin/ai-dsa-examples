/*
In the Heap Sort algorithm, we first build a max-heap from the given array. A max-heap is a complete binary tree where the value of each node is greater than or equal to the values of its children. Once the max-heap is built, we repeatedly extract the maximum element (root) from the heap and place it at the end of the array. This process is repeated until the entire array is sorted in ascending order.

The HeapSort function takes an array as input and modifies it in-place to sort it in ascending order using Heap Sort. In the example, the array [5, 2, 4, 6, 1, 3] is sorted to [1, 2, 3, 4, 5, 6] using Heap Sort.
*/

package main

import "fmt"

func main() {
    arr := []int{5, 2, 4, 6, 1, 3}
    HeapSort(arr)
    fmt.Println(arr) // Output: [1 2 3 4 5 6]
}

func HeapSort(arr []int) {
    n := len(arr)

    // Build a max-heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    // Extract elements from the heap in descending order
    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0] // Swap the root (maximum element) with the last element
        heapify(arr, i, 0)              // Heapify the reduced heap
    }
}

func heapify(arr []int, n, i int) {
    largest := i       // Initialize the largest element as the root
    left := 2*i + 1    // Left child
    right := 2*i + 2   // Right child

    // If the left child is larger than the root
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    // If the right child is larger than the largest so far
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    // If the largest element is not the root
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i] // Swap the root with the largest element
        heapify(arr, n, largest)                    // Recursively heapify the affected subtree
    }
}