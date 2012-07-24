package mobile

import (
	"testing"
)

var mobileNums = []struct {
	mobileNum string
	b         bool
}{
	/* 13 */
	{"13000000000", true},
	{"13100000000", true},
	{"13200000000", true},
	{"13300000000", true},
	{"13400000000", true},
	{"13500000000", true},
	{"13600000000", true},
	{"13700000000", true},
	{"13800000000", true},
	{"13900000000", true},
	/* 14 */
	{"14500000000", true},
	{"14700000000", true},
	/* 15 */
	{"15000000000", true},
	{"15100000000", true},
	{"15200000000", true},
	{"15300000000", true},
	{"15500000000", true},
	{"15600000000", true},
	{"15700000000", true},
	{"15800000000", true},
	{"15900000000", true},
	/* 18 */
	{"18000000000", true},
	{"18700000000", true},
	{"18800000000", true},
	{"18900000000", true},

	/* false */
	{"1340000000", false},
	{"15400000000", false},
	{"", false},
	{"123", false},
	{"12300000000", false},
	{"14000000000", false},
	{"14100000000", false},
	{"14200000000", false},
	{"14300000000", false},
	{"14400000000", false},
	{"14600000000", false},
	{"14800000000", false},
	{"14900000000", false},
	{"1490000000a", false},
	{"1490a000000", false},
}

func TestCheck(t *testing.T) {
	for _, v := range mobileNums {
		b := Validate(v.mobileNum)
		if b != v.b {
			t.Error("b!=v.b",v.mobileNum, b, v.b)
		}
	}
}
