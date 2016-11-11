package quicksort

import (
	"fmt"
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
		actual := Sort(tt.arr, byFirstElement)
		if !isEqual(actual, tt.expected) {
			t.Errorf("Sort(\"%v\") expected %v, actual %v", tt.arr, tt.expected, actual)
		}
	}
}

func TestSortByLastElement(t *testing.T) {
	for _, tt := range SortTests {
		actual := Sort(tt.arr, byLastElement)
		if !isEqual(actual, tt.expected) {
			t.Errorf("Sort(\"%v\") expected %v, actual %v", tt.arr, tt.expected, actual)
		}
	}
}

func TestSortByMedianOfThree(t *testing.T) {
	for _, tt := range SortTests {
		actual := Sort(tt.arr, byMedianOfThree)
		if !isEqual(actual, tt.expected) {
			t.Errorf("Sort(\"%v\") expected %v, actual %v", tt.arr, tt.expected, actual)
		}
	}
}

/*
func TestSortFromFile(t *testing.T) {
	var act []int
	fmt.Println("10.txt:")
	act = SortFromFile("test/10.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", act)
	act = SortFromFile("test/10.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", act)
	act = SortFromFile("test/10.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", act)
	fmt.Println("1000.txt:")
	act = SortFromFile("test/1000.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", act)
	act = SortFromFile("test/1000.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", act)
	act = SortFromFile("test/1000.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", act)
}
*/

func TestCountSortFromFile(t *testing.T) {
	var count int

	fmt.Println("10.txt:")
	_, count = CountSortFromFile("test/10.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count)
	_, count = CountSortFromFile("test/10.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count)
	_, count = CountSortFromFile("test/10.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count)

	fmt.Println("10_sorted.txt:")
	_, count = CountSortFromFile("test/10_sorted.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count)
	_, count = CountSortFromFile("test/10_sorted.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count)
	_, count = CountSortFromFile("test/10_sorted.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count)

	fmt.Println("100.txt:")
	_, count = CountSortFromFile("test/100.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count)
	_, count = CountSortFromFile("test/100.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count)
	_, count = CountSortFromFile("test/100.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count)

	fmt.Println("1000.txt:")
	_, count = CountSortFromFile("test/1000.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count)
	_, count = CountSortFromFile("test/1000.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count)
	_, count = CountSortFromFile("test/1000.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count)

	fmt.Println("QuickSort.txt:")
	_, count = CountSortFromFile("test/QuickSort.txt", byFirstElement)
	fmt.Printf("byFirstElement: %v\n", count)
	_, count = CountSortFromFile("test/QuickSort.txt", byLastElement)
	fmt.Printf("byLastElement: %v\n", count)
	_, count = CountSortFromFile("test/QuickSort.txt", byMedianOfThree)
	fmt.Printf("byMedianOfThree: %v\n", count)
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
