package coins

import "errors"

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// Coin is a coin with a value and quantity
type Coin struct {
	value    int
	Quantity int
}

// Coins is a slice of Coins
type Coins map[int]Coin

// Minimum takes an amount and a set of Coins (quantities ignored) and returns
// the Coins with minimum total quantity that sum to amount.
func Minimum(amount int, coins Coins) (Coins, int) {
	m := make(map[int]int)
	s := make(map[int]int)

	m[0] = 0
	for n := 1; n <= amount; n++ {
		var coin int
		min := maxInt
		for i, c := range coins {
			if c.value <= n && n-c.value >= 0 && (1+m[n-c.value]) < min {
				if m[n-c.value] != maxInt {
					min = 1 + m[n-c.value]
				}
				coin = i
			}
		}
		m[n] = min
		s[n] = coin
	}

	for _, c := range coins {
		c.Quantity = 0
	}

	count := 0
	for amount > 0 {
		i := s[amount]
		inc := coins[i]
		inc.Quantity++
		count++
		amount -= coins[i].value
	}

	return coins, count
}

// Combinations takes an amount and a set of Coins (quantities ignored) and
// returns the number of unique combinations of those coins that sum to amount.
func Combinations(amount int, coins Coins) (int, error) {
	for _, c := range coins {
		c.Quantity = 0
	}

	C := make(map[int][]Coins)
	C[0] = append(C[0], coins.duplicate())

	// n
	for n := 0; n <= amount; n++ {
		// n^c
		for _, combo := range C[n] {
		CombinationLoop:
			// c
			for v, c := range coins {
				for _, cc := range combo {
					if cc.value > c.value && cc.Quantity > 0 {
						continue CombinationLoop
					}
				}
				if n+c.value <= amount {
					nextCombo := combo.duplicate()
					inc := nextCombo[v]
					inc.Quantity++
					nextCombo[v] = inc
					C[n+c.value] = append(C[n+c.value], nextCombo)
				}
			}
		}
	}

	combos, ok := C[amount]
	if !ok {
		return 0, errors.New("No defined solution")
	}
	return len(combos), nil
}

func (coins Coins) duplicate() Coins {
	c := Coins{}
	for val, coin := range coins {
		c[val] = Coin{coin.value, coin.Quantity}
	}
	return c
}
