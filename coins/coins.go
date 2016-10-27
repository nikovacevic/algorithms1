package coins

import "sort"

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// Coin is a coin with a value and quantity
type Coin struct {
	value    int
	Quantity int
}

// Coins is a slice of Coins
type Coins []Coin

// ByValue implements sort.Interface for Coins based on Coin.value
type ByValue Coins

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

// ByQuantity implements sort.Interface for Coins based on Coin.Quantity
type ByQuantity Coins

func (a ByQuantity) Len() int           { return len(a) }
func (a ByQuantity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByQuantity) Less(i, j int) bool { return a[i].Quantity < a[j].Quantity }

// Equals tests equality of two Coin instances
func (a Coin) Equals(b Coin) bool {
	return a.value == b.value && a.Quantity == b.Quantity
}

// Equals tests equality of two sets of Coins. Order matters!
func (a Coins) Equals(b Coins) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !(a[i].value == b[i].value && a[i].Quantity == b[i].Quantity) {
			return false
		}
	}
	return true
}

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
		coins[i].Quantity++
		count++
		amount -= coins[i].value
	}

	return coins, count
}

// Combinations takes an amount and a slice of Coins and returns the number of
// unique combinations of those coins that sum the given amount.
func Combinations(amount int, coins Coins) int {
	// Sort coins by value
	sort.Sort(ByValue(coins))
	// Map of amount to combination count
	C := make(map[int]int)
	C[0] = 1
	// Iterate over all coins, rom least to greatest value, checking
	// combination counts for all possible amounts (C[amt]) for the given
	// coin (i.e. those greater than or equal to the coin's value) and less
	// than the given amount. For any C[amt] and coin, C[amt-coin.value]
	// represents the count of coin combinations necessary to get to the
	// amount one coin.value less than amount (i.e. the previous step).
	for _, coin := range coins {
		for amt := coin.value; amt <= amount; amt++ {
			C[amt] += C[amt-coin.value]
		}
	}
	// Return the combination count at the given amount
	return C[amount]
}
