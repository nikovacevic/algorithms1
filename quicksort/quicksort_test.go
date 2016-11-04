package quicksort

import (
	"fmt"
	"math/big"
	"testing"
)

var SortTests = []struct {
	arr      []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{3, 2, 1, 5, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
	{[]int{3, 3, 3, 2, 1, 5, 4, 6, 21, -10, -10}, []int{-10, -10, 1, 2, 3, 3, 3, 4, 5, 6, 21}},
}

func TestSortByFirstElement(t *testing.T) {
	for _, tt := range SortTests {
		actual, _ := Sort(tt.arr, byFirstElement, big.Int{})
		if !isEqual(actual, tt.expected) {
			t.Errorf("Sort(\"%v\") expected %v, actual %v", tt.arr, tt.expected, actual)
		}
	}
}

func TestSortByLastElement(t *testing.T) {
	for _, tt := range SortTests {
		actual, _ := Sort(tt.arr, byLastElement, big.Int{})
		if !isEqual(actual, tt.expected) {
			t.Errorf("Sort(\"%v\") expected %v, actual %v", tt.arr, tt.expected, actual)
		}
	}
}

func TestSortByMedianOfThree(t *testing.T) {
	for _, tt := range SortTests {
		actual, _ := Sort(tt.arr, byMedianOfThree, big.Int{})
		if !isEqual(actual, tt.expected) {
			t.Errorf("Sort(\"%v\") expected %v, actual %v", tt.arr, tt.expected, actual)
		}
	}
}

func TestSortFromFile(t *testing.T) {
	count := big.Int{}
	fmt.Println("10.txt:")
	_, count = SortFromFile("test/10.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/10.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/10.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count.Text(10))
	fmt.Println("100.txt:")
	_, count = SortFromFile("test/100.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/100.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/100.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count.Text(10))
	fmt.Println("1000.txt:")
	_, count = SortFromFile("test/1000.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/1000.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/1000.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count.Text(10))
	fmt.Println("QuickSort.txt:")
	_, count = SortFromFile("test/QuickSort.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/QuickSort.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count.Text(10))
	_, count = SortFromFile("test/QuickSort.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count.Text(10))
}

func isEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
