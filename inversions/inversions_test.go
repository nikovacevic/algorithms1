package inversions

import "testing"

var CountTests = []struct {
	arr []int
	exp int
}{
	{[]int{1}, 0},
	{[]int{1, 2, 3}, 0},
	{[]int{3, 2, 1}, 3},
	{[]int{1, 2, 4, 3, 5, 7, 8, 6}, 3},
}

func TestCount(t *testing.T) {
	for _, tt := range CountTests {
		act := Count(tt.arr)
		if act != tt.exp {
			t.Errorf("Count(\"%v\") expected %v, actual %v", tt.arr, tt.exp, act)
		}
	}
}

func TestCountFromFile(t *testing.T) {
	exp := 2407905288
	act, err := CountFromFile("test/IntegerArray.txt")
	if err != nil {
		t.Errorf("CountFromFile(\"%v\") gave error %v", "test/IntegerArray.txt", err)
	}
	if act != exp {
		t.Errorf("CountFromFile(\"%v\") expected %v, actual %v", "test/IntegerArray.txt", exp, act)
	}
}
