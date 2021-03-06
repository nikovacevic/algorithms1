package quicksort

// pivotFunc is a function that selects a pivot element index
type pivotFunc func([]int) int

// Sort sorts a slice of integers in ascending order
func Sort(a []int, pf pivotFunc) []int {
	// Base case: an array of length 1 is always sorted
	if len(a) <= 1 {
		return a
	}
	p, pivot := selectPivot(pf, a)
	a, i := Partition(a, p)
	al := Sort(a[:i-1], pf)
	ar := Sort(a[i:], pf)
	a = append(append(al, pivot), ar...)
	return a
}

// Partition rearranges the elements of a such that elements 0 through i are
// less than or equal to a[p] and elements i+1 an on are greater than a[p]
func Partition(a []int, p int) ([]int, int) {
	pivot := a[p]
	// Swap pivot with first element
	a[0], a[p] = a[p], a[0]
	var i, j int
	for i, j = 1, 1; j < len(a); j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	// Swap pivot with last element smaller than pivot
	a[0], a[i-1] = a[i-1], a[0]
	return a, i
}

// selectPivot selects a pivot element in a based on the given pivot function,
// returning (pivot index, pivot value)
func selectPivot(pf pivotFunc, a []int) (int, int) {
	p := pf(a)
	return p, a[p]
}

// byFirstElement is a pivotFunc that always selects the first element
func byFirstElement(a []int) int {
	return 0
}

// byLastElement is a pivotFunc that always selects the last element
func byLastElement(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return len(a) - 1
}

// byMedianOfThree is a pivotFunc that selects the median of the first, last
// and middle element
func byMedianOfThree(a []int) int {
	if len(a) <= 2 {
		return 0
	}
	i := []int{0, len(a) / 2, len(a) - 1}
	s := []int{a[0], a[len(a)/2], a[len(a)-1]}
	for x := range s {
		y := x
		for y > 0 && s[y-1] > s[y] {
			i[y], i[y-1] = i[y-1], i[y]
			s[y], s[y-1] = s[y-1], s[y]
			y--
		}
	}
	return i[1]
}
