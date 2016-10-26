package coins

import "testing"

var MinimumTests = []struct {
	amount int
	coins  Coins
	exp    int
}{
	{7, Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, 3},
	{9, Coins{2: Coin{2, 0}, 3: Coin{3, 0}}, 3},
	{10, Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, 4},
	{100, Coins{2: Coin{2, 0}, 3: Coin{3, 0}, 10: Coin{10, 0}, 46: Coin{46, 0}}, 5},
}

func TestMinimum(t *testing.T) {
	for _, tt := range MinimumTests {
		_, act := Minimum(tt.amount, tt.coins)
		if act != tt.exp {
			t.Errorf("Minimum(%v, %v) expected %v, actual %v", tt.amount, tt.coins, tt.exp, act)
		}
	}
}

var CombinationsTests = []struct {
	amount int
	coins  Coins
	exp    int
}{
	{amount: 1, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, exp: 1},
	{amount: 7, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, exp: 8},
	{amount: 20, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, exp: 44},
	{amount: 30, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, exp: 91},
	{amount: 40, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, exp: 154},
	{amount: 80, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}}, exp: 574},
	{amount: 80, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}, 4: Coin{4, 0}, 5: Coin{5, 0}, 6: Coin{6, 0}}, exp: 69624},
	//{amount: 80, coins: Coins{1: Coin{1, 0}, 2: Coin{2, 0}, 3: Coin{3, 0}, 4: Coin{4, 0}, 5: Coin{5, 0}, 6: Coin{6, 0}, 7: Coin{7, 0}, 8: Coin{8, 0}}, exp: 69624},
	//{amount: 300, coins: Coins{1: Coin{1, 0}, 5: Coin{5, 0}, 10: Coin{10, 0}, 20: Coin{20, 0}}, exp: 5456},
}

func TestCombinations(t *testing.T) {
	for _, tt := range CombinationsTests {
		if act, err := Combinations(tt.amount, tt.coins); err != nil {
			t.Errorf("Combinations(%v, %v) expected %v, actual %v", tt.amount, tt.coins, tt.exp, act)
		} else {
			if act != tt.exp {
				t.Errorf("Combinations(%v, %v) expected %v, actual %v", tt.amount, tt.coins, tt.exp, act)
			}
		}
	}
}
