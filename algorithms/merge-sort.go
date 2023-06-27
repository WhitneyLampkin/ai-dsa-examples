/*
In the Merge Sort algorithm, we divide the array into two halves recursively until we reach subarrays of size 1 or 0. Then, we merge these subarrays back together in sorted order. This process is repeated until the entire array is sorted.

The MergeSort function takes an array as input and returns a new sorted array using Merge Sort. In the example, the array [5, 2, 4, 6, 1, 3] is sorted to [1, 2, 3, 4, 5, 6] using Merge Sort. The original array remains unchanged, and a new sorted array is returned.
*/

package main

import "fmt"

func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // Divide the array into two halves
    mid := len(arr) / 2
    left := MergeSort(arr[:mid])
    right := MergeSort(arr[mid:])

    // Merge the two sorted halves
    return merge(left, right)
}

func merge(left, right []int) []int {
    merged := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    // Compare elements from the left and right halves and merge them in sorted order
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            merged = append(merged, left[i])
            i++
        } else {
            merged = append(merged, right[j])
            j++
        }
    }

    // Add the remaining elements from the left or right half (if any)
    merged = append(merged, left[i:]...)
    merged = append(merged, right[j:]...)

    return merged
}

func main() {
    arr := []int{5, 2, 4, 6, 1, 3}
    sortedArr := MergeSort(arr)
    fmt.Println(sortedArr) // Output: [1 2 3 4 5 6]
}