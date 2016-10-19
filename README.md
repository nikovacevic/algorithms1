# Algorithms 1

[Coursera Algorithms I](https://www.coursera.org/learn/algorithms-divide-conquer/home/welcome), implemented in Go.

## Week 1: Divide and Conquer

Covers merge sort, Karatsuba's integer multiplication, basic O-notation analysis.

### [Mergesort](https://github.com/nikovacevic/algorithms1/blob/master/mergesort/mergesort.go)

`Sort` returns the array (`O(1)`) in the base case. It recursively calls itself twice for each non-base-case call. This can happen at most `log n` times before each recursive call reached the base case.

`merge` that returns an array of length `l` will run in `O(l)` time, given that one instance of constant-factor comparisons and element copying is required for each element. We ignore constant factors to get `O(l)`.

Given that `merge` will be called at most `log n` times, and `n` (number of elements in the original array) is an upper-bound on `l`, we can say that the entire process has an asymtotic upper-bound of `O(n log n)`.

### [Karatsuba multiplication](https://github.com/nikovacevic/algorithms1/blob/master/multiplication/multiplication.go)

`Karatsuba` acts as a wrapper function, executing some low-factor (O(`n`) or O(`1`)) helper functions to pre-process the input to `karatsuba`, which does the real work.

`karatsuba` can be modeled with the following recurrence relation:
```
T(n) = 3*T(n/2) + 2A(n/2) + A(n) + S(n)
```
where `A(n)` defines efficiency of `add` and `S(n)` for `sub`. We want to isolate the non-recursive bit, then analyze that:
```
T(n) = 3*T(n/2) + f(n)
where f(n) = 2A(n/2) + A(n) + S(n)
and A(n) = Θ(n)
and S(n) = Θ(n)
therefore, f(n) = Θ(n)
```
Using the [Master theorem](https://en.wikipedia.org/wiki/Master_theorem), we state the following:
```
T(n) = 3*T(n/2) + f(n), so a = 3, b = 2

By the Master method, if there exists c such that f(n)=O(n^c) where c<log_b(a), then T(n)=Θ(n^log_b(a))

We know f(n) = Θ(n), so choose c=1
Then, f(n)=O(n) where 1<log_2(3), and therefore
T(n)=Θ(n^log_2(3))
```
