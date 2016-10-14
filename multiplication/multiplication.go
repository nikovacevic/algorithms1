package multiplication

import (
	"math"
	"strconv"
)

// Karatsuba returns the product of a and b, which should be positive integers
// represented as strings (e.g. 123 as "123").
func Karatsuba(a string, b string) string {
	// Make sure string lengths are equal
	evenPad(&a, &b)
	// Make sure string lengths are powers of 2
	a = padToPowerOfTwo(a)
	b = padToPowerOfTwo(b)
	return unpad(karatsuba(a, b))
}

// karatsuba implements Karatsuba's recursive algorithm, assuming a and b
// have the same length, the length being a power of 2.
func karatsuba(a string, b string) string {
	if len(a) == 1 {
		// Convert to integers and multiply
		aint, _ := strconv.Atoi(a)
		bint, _ := strconv.Atoi(b)
		return strconv.Itoa(aint * bint)
	}

	n := len(a)
	h := n / 2
	al := a[:h]
	ar := a[h:]
	bl := b[:h]
	br := b[h:]
	x := karatsuba(al, bl)
	y := add(karatsuba(al, br), karatsuba(bl, ar))
	z := karatsuba(ar, br)

	return add(add(shift(x, n), shift(y, h)), z)
}

func add(a string, b string) string {
	evenPad(&a, &b)

	sum := ""
	c := 0
	for i := len(a) - 1; i >= 0; i-- {
		aint, _ := strconv.Atoi(string(a[i]))
		bint, _ := strconv.Atoi(string(b[i]))
		s := aint + bint + c
		if s >= 10 {
			c = 1
			s -= 10
		} else {
			c = 0
		}
		sum = strconv.Itoa(s) + sum
	}

	if c > 0 {
		sum = strconv.Itoa(c) + sum
	}

	return sum
}

func sub(a string, b string) string {
	evenPad(&a, &b)

	diff := ""
	c := false
	for i := len(a) - 1; i >= 0; i-- {
		aint, _ := strconv.Atoi(string(a[i]))
		bint, _ := strconv.Atoi(string(b[i]))

		if c {
			aint--
		}

		if bint > aint {
			aint += 10
			c = true
		} else {
			c = false
		}

		d := aint - bint

		diff = strconv.Itoa(d) + diff
	}

	return unpad(diff)
}

func evenPad(a *string, b *string) {
	if len(*a) != len(*b) {
		for len(*a) < len(*b) {
			*a = "0" + *a
		}
		for len(*b) < len(*a) {
			*b = "0" + *b
		}
	}
}

func padToPowerOfTwo(a string) string {
	l := len(a)
	p := 1
	e := 0
	for ; l > p; e++ {
		p = int(math.Pow(2, float64(e)))
	}
	for ; l < p; l++ {
		a = "0" + a
	}
	return a
}

func unpad(a string) string {
	var r string
	for i, l := range a {
		if l != '0' {
			r = a[i:]
			break
		}
	}
	if r == "" {
		r = "0"
	}
	return r
}

func shift(a string, n int) string {
	for i := 0; i < n; i++ {
		a = a + "0"
	}
	return a
}
