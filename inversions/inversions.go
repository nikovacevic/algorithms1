package inversions

import (
	"bufio"
	"os"
	"strconv"
)

// Count counts the number of inversions in the array of integers
func Count(arr []int) int {
	inv := 0
	_, inv = sortInv(arr)
	return inv
}

// sortInv leverages mergesort to sort the given array, counting and returning
// the inversions as well.
func sortInv(arr []int) ([]int, int) {
	// Base case: an array of length 1 is always sorted
	if len(arr) <= 1 {
		return arr, 0
	}
	// Merge the recursively sorted halves of the array
	sarr1, i1 := sortInv(arr[:len(arr)/2])
	sarr2, i2 := sortInv(arr[len(arr)/2:])
	sarr, i := mergeInv(sarr1, sarr2)
	return sarr, (i + i1 + i2)
}

// merge is a utility function used by Sort to merge two arrays while
// maintaining ascenting order
func mergeInv(arr1 []int, arr2 []int) ([]int, int) {
	// Merged array
	var arr []int
	// Track number of inversions
	inv := 0
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
			inv += len(arr1) - i1
			i2++
		}
	}
	return arr, inv
}

// CountFromFile takes the full path to a text file representing an array
// of integers, with one number per line, and returns the number of
// inversions in that file.
func CountFromFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var arr []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		arr = append(arr, i)
	}
	return Count(arr), nil
}
