package coins

import (
	"sort"
	"testing"
)

var MinimumTests = []struct {
	amount int
	coins  Coins
	exp    int
}{
	{7, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}, 3},
	{9, Coins{Coin{2, 0}, Coin{3, 0}}, 3},
	{10, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}, 4},
	{100, Coins{Coin{2, 0}, Coin{3, 0}, Coin{10, 0}, Coin{46, 0}}, 5},
}

func TestMinimum(t *testing.T) {
	for _, tt := range MinimumTests {
		_, act := Minimum(tt.amount, tt.coins)
		if act != tt.exp {
			t.Errorf("Minimum(%v, %v) expected %v, actual %v", tt.amount, tt.coins, tt.exp, act)
		}
	}
}

var SortByValueTests = []struct {
	coins Coins
	exp   Coins
}{
	{Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}},
	{Coins{Coin{3, 0}, Coin{2, 0}, Coin{1, 0}}, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}},
	{Coins{Coin{3, 0}, Coin{1, 0}, Coin{3, 0}}, Coins{Coin{1, 0}, Coin{3, 0}, Coin{3, 0}}},
}

func TestSortByValue(t *testing.T) {
	for _, tt := range SortByValueTests {
		sort.Sort(ByValue(tt.coins))
		if !tt.coins.Equals(tt.exp) {
			t.Errorf("sort.Sort(ByValue(%v)) expected %v, actual %v", tt.coins, tt.exp, tt.coins)
		}
	}
}

var CombinationsTests = []struct {
	amount int
	coins  Coins
	exp    int
}{
	{0, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}, 1},
	{1, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}, 1},
	{7, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}}, 8},
	{80, Coins{Coin{1, 0}, Coin{2, 0}, Coin{3, 0}, Coin{4, 0}, Coin{5, 0}, Coin{6, 0}}, 69624},
	{300, Coins{Coin{1, 0}, Coin{5, 0}, Coin{10, 0}, Coin{20, 0}}, 5456},
	{3000, Coins{Coin{1, 0}, Coin{5, 0}, Coin{10, 0}, Coin{20, 0}}, 4590551},
	{3000, Coins{Coin{10, 0}, Coin{1, 0}, Coin{5, 0}, Coin{20, 0}}, 4590551},
}

func TestCombinations(t *testing.T) {
	for _, tt := range CombinationsTests {
		act := Combinations(tt.amount, tt.coins)
		if act != tt.exp {
			t.Errorf("Combinations(%v, %v) expected %v, actual %v", tt.amount, tt.coins, tt.exp, act)
		}
	}
}
