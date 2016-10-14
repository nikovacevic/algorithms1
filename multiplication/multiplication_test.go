package multiplication

import "testing"

var KaratsubaTests = []struct {
	a   string
	b   string
	exp string
}{
	{"10", "10", "100"},
	{"12", "43", "516"},
	{"1234", "4321", "5332114"},
	{"123", "1234", "151782"},
	{"3141592653589793238462643383279502884197169399375105820974944592", "2718281828459045235360287471352662497757247093699959574966967627", "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"},
}

func TestKaratsuba(t *testing.T) {
	for _, tt := range KaratsubaTests {
		act := Karatsuba(tt.a, tt.b)
		if act != tt.exp {
			t.Errorf("Karatsuba(\"%s\", \"%s\") expected %s, actual %s", tt.a, tt.b, tt.exp, act)
		}
	}
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

var padToPowerOfTwoTests = []struct {
	a   string
	exp string
}{
	{"10", "10"},
	{"100", "0100"},
	{"123456", "00123456"},
}

func TestPadToPowerOfTwo(t *testing.T) {
	for _, tt := range padToPowerOfTwoTests {
		act := padToPowerOfTwo(tt.a)
		if act != tt.exp {
			t.Errorf("padToPowerOfTwo(\"%v\") expected %v, actual %v", tt.a, tt.exp, act)
		}
	}
}

var shiftTests = []struct {
	a   string
	n   int
	exp string
}{
	{"1", 1, "10"},
	{"123", 3, "123000"},
	{"1", 0, "1"},
}

func TestShift(t *testing.T) {
	for _, tt := range shiftTests {
		act := shift(tt.a, tt.n)
		if act != tt.exp {
			t.Errorf("shift(\"%v\", \"%v\") expected %v, actual %v", tt.a, tt.n, tt.exp, act)
		}
	}
}
