package multiplication

import "testing"

func TestKaratsuba(t *testing.T) {

}

var addTests = []struct {
	a   string
	b   string
	exp string
}{
	{"0", "0", "0"},
	{"10", "1", "11"},
	{"999", "9", "1008"},
}

func TestAdd(t *testing.T) {
	for _, tt := range addTests {
		act := add(tt.a, tt.b)
		if act != tt.exp {
			t.Errorf("add(\"%v\", \"%v\") expected %v, actual %v", tt.a, tt.b, tt.exp, act)
		}
	}
}

var subTests = []struct {
	a   string
	b   string
	exp string
}{
	{"0", "0", "0"},
	{"11", "10", "1"},
	{"1008", "999", "9"},
	{"912381248619", "823981723741", "88399524878"},
}

func TestSub(t *testing.T) {
	for _, tt := range subTests {
		act := sub(tt.a, tt.b)
		if act != tt.exp {
			t.Errorf("sub(\"%v\", \"%v\") expected %v, actual %v", tt.a, tt.b, tt.exp, act)
		}
	}
}

var evenPadTests = []struct {
	a    string
	b    string
	expa string
	expb string
}{
	{"0", "0", "0", "0"},
	{"01", "10", "01", "10"},
	{"0", "10", "00", "10"},
	{"000", "10", "000", "010"},
}

func TestEvenPad(t *testing.T) {
	for _, tt := range evenPadTests {
		befa, befb := tt.a, tt.b
		evenPad(&tt.a, &tt.b)
		if tt.a != tt.expa || tt.b != tt.expb {
			t.Errorf("evenPad(\"%v\", \"%v\") expected (%v, %v), actual (%v, %v)", befa, befb, tt.expa, tt.expb, tt.a, tt.b)
		}
	}
}
