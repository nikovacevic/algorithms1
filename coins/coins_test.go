package coins

import "testing"

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
