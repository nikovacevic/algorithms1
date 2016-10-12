# Algorithms 1

[Coursera Algorithms I](https://www.coursera.org/learn/algorithms-divide-conquer/home/welcome), implemented in Go.

## Week 1: Divide and Conquer

Covers merge sort, Karatsuba's integer multiplication, basic O-notation analysis.

### [Mergesort](https://github.com/nikovacevic/algorithms1/blob/master/mergesort/mergesort.go)

`Sort` returns the array (`O(1)`) in the base case. It recursively calls itself twice for each non-base-case call. This can happen at most `log n` times before each recursive call reached the base case.

`merge` that returns an array of length `l` will run in `O(l)` time, given that one instance of constant-factor comparisons and element copying is required for each element. We ignore constant factors to get `O(l)`.

Given that `merge` will be called at most `log n` times, and `n` (number of elements in the original array) is an upper-bound on `l`, we can say that the entire process has an asymtotic upper-bound of `O(n log n)`.
