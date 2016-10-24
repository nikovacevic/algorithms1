package coins

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// Coin is a coin with a value and quantity
type Coin struct {
	value    int
	quantity int
}

// Coins is a slice of Coins
type Coins []Coin

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
		c.quantity = 0
	}

	count := 0
	for amount > 0 {
		i := s[amount]
		coins[i].quantity++
		count++
		amount -= coins[i].value
	}

	return coins, count
}

// Combinations takes an amount and a set of Coins (quantities ignored) and
// returns the number of unique combinations of those coins that sum to amount.
func Combinations(amount int, c Coins) int {
	return 0
}
