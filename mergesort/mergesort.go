package mergesort

import "fmt"

// Sort sorts a slice of integers in ascending order, returning the result
func Sort(arr []int) []int {
	// Base case: an array of length 1 is always sorted
	if len(arr) == 1 {
		return arr
	}
	// Merge the recursively sorted halves of the array
	return merge(Sort(arr[:len(arr)/2]), Sort(arr[len(arr)/2:]))
}

// merge is a utility function used by Sort to merge two arrays while
// maintaining ascenting order
func merge(arr1 []int, arr2 []int) []int {
	fmt.Printf("Merge %v with %v => ", arr1, arr2)
	// Merged array
	var arr []int
	// Use two indices to track progress through each array
	i1, i2 := 0, 0
	for i1 < len(arr1) || i2 < len(arr2) {
		if i1 == len(arr1) {
			arr = append(arr, arr2[i2])
			i2++
		} else if i2 == len(arr2) {
			arr = append(arr, arr1[i1])
			i1++
		} else if arr1[i1] < arr2[i2] {
			arr = append(arr, arr1[i1])
			i1++
		} else {
			arr = append(arr, arr2[i2])
			i2++
		}
	}
	fmt.Printf("%v\n", arr)
	return arr
}
