/*
In the Quick Sort algorithm, we select an element from the array as the pivot and partition the other elements into two subarrays, according to whether they are less than or greater than the pivot. We then recursively apply the same process to the left and right subarrays until the entire array is sorted.

The QuickSort function takes an array, along with the low and high indices, as input and modifies it in-place to sort it in ascending order using Quick Sort. In the example, the array [5, 2, 4, 6, 1, 3] is sorted to [1, 2, 3, 4, 5, 6] using Quick Sort.
*/

package main

import "fmt"

func main() {
    arr := []int{5, 2, 4, 6, 1, 3}
    QuickSort(arr, 0, len(arr)-1)
    fmt.Println(arr) // Output: [1 2 3 4 5 6]
}

func QuickSort(arr []int, low, high int) {
    if low < high {
        // Partition the array and get the pivot index
        pivotIndex := partition(arr, low, high)

        // Recursively sort the left and right subarrays
        QuickSort(arr, low, pivotIndex-1)
        QuickSort(arr, pivotIndex+1, high)
    }
}

func partition(arr []int, low, high int) int {
    // Choose the rightmost element as the pivot
    pivot := arr[high]

    // Index of the smaller element
    i := low - 1

    for j := low; j < high; j++ {
        // If the current element is smaller than or equal to the pivot
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i] // Swap arr[i] and arr[j]
        }
    }

    // Place the pivot element at its correct position
    arr[i+1], arr[high] = arr[high], arr[i+1]

    return i + 1
}