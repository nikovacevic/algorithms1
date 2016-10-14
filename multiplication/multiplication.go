package multiplication

import (
	"fmt"
	"strconv"
)

// Karatsuba returns the product of a and b
// TODO more
func Karatsuba(a string, b string) (string, error) {
	// Make sure string lengths are equal
	evenPad(&a, &b)
	// Make sure string lengths are powers of 2
	// TODO
	r, err := karatsuba(a, b)
	if err != nil {
		return "", err
	}
	return r, nil
}

// karatsuba implements Karatsuba's recursive algorithm, assuming a and b
// have the same length.
func karatsuba(a string, b string) (string, error) {
	if len(a) == 1 {
		// Convert to integers and multiply
		aint, erra := strconv.Atoi(a)
		if erra != nil {
			return "", erra
		}
		bint, errb := strconv.Atoi(b)
		if errb != nil {
			return "", errb
		}
		return string(aint * bint), nil
	}

	h := len(a) / 2
	al := a[:h]
	ar := a[h:]
	bl := b[:h]
	br := b[h:]

	x, errx := karatsuba(al, bl)
	if errx != nil {
		return "", errx
	}
	y, erry := karatsuba(ar, br)
	if erry != nil {
		return "", erry
	}
	z, errz := karatsuba(add(al, ar), add(bl, br))
	if errz != nil {
		return "", errz
	}

	// (10^n)(a)(c) + (10^(n/2))((a*d)+(b*c))+ (b)(d)
	// TODO
	return x + y + z, nil
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

	fmt.Println(diff)

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
