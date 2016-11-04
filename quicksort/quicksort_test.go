package quicksort

import "testing"

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
