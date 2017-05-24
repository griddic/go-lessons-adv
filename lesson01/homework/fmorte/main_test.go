package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OnePiece(t *testing.T) {
	assert.Equal(t, "штука", Format(1), "1 штука")
}

func Test_OneDigitShtuki(t *testing.T) {
	assert.Equal(t, "штуки", Format(2), "2 штуки")
	assert.Equal(t, "штуки", Format(3), "3 штуки")
	assert.Equal(t, "штуки", Format(4), "4 штуки")
}

func Test_OneDigitShtuk(t *testing.T) {
	assert.Equal(t, "штук", Format(5), "5 штук")
	assert.Equal(t, "штук", Format(6), "6 штук")
	assert.Equal(t, "штук", Format(7), "7 штук")
	assert.Equal(t, "штук", Format(8), "8 штук")
	assert.Equal(t, "штук", Format(9), "9 штук")
}

func Test_Teens(t *testing.T) {
	assert.Equal(t, "штук", Format(11), "11 штук")
	assert.Equal(t, "штук", Format(12), "12 штук")
	assert.Equal(t, "штук", Format(13), "13 штук")
	assert.Equal(t, "штук", Format(14), "14 штук")
	assert.Equal(t, "штук", Format(15), "15 штук")
	assert.Equal(t, "штук", Format(16), "16 штук")
	assert.Equal(t, "штук", Format(17), "17 штук")
	assert.Equal(t, "штук", Format(18), "18 штук")
	assert.Equal(t, "штук", Format(19), "19 штук")
}

func Test_LowerThan100(t *testing.T) {
	assert.Equal(t, "штук", Format(27), "27 штук")
	assert.Equal(t, "штук", Format(45), "45 штук")
	assert.Equal(t, "штука", Format(61), "61 штука")
	assert.Equal(t, "штуки", Format(32), "32 штуки")
	assert.Equal(t, "штуки", Format(93), "93 штуки")
	assert.Equal(t, "штуки", Format(84), "84 штуки")
}

func Test_TeensAfter100(t *testing.T) {
	assert.Equal(t, "штук", Format(111), "111 штук")
	assert.Equal(t, "штук", Format(112), "112 штук")
	assert.Equal(t, "штук", Format(113), "113 штук")
	assert.Equal(t, "штук", Format(114), "114 штук")
	assert.Equal(t, "штук", Format(115), "115 штук")
	assert.Equal(t, "штук", Format(116), "116 штук")
	assert.Equal(t, "штук", Format(117), "117 штук")
	assert.Equal(t, "штук", Format(118), "118 штук")
	assert.Equal(t, "штук", Format(119), "119 штук")
}

func Test_TeensAfter1000(t *testing.T) {
	assert.Equal(t, "штук", Format(1011), "1011 штук")
	assert.Equal(t, "штук", Format(1012), "1012 штук")
	assert.Equal(t, "штук", Format(1013), "1013 штук")
	assert.Equal(t, "штук", Format(1014), "1014 штук")
	assert.Equal(t, "штук", Format(1015), "1015 штук")
	assert.Equal(t, "штук", Format(1016), "1016 штук")
	assert.Equal(t, "штук", Format(1017), "1017 штук")
	assert.Equal(t, "штук", Format(1018), "1018 штук")
	assert.Equal(t, "штук", Format(1019), "1019 штук")
}

func Test_NearInt32UpperLimit(t *testing.T) {
	assert.Equal(t, "штука", Format(2147483641), "2147483641 штука")
	assert.Equal(t, "штуки", Format(2147483644), "2147483644 штуки")
	assert.Equal(t, "штук", Format(2147483647), "2147483647 штук")
	assert.Equal(t, "штук", Format(2147483611), "2147483611 штук")
}

func Test_Zero(t *testing.T) {
	assert.Equal(t, "штук", Format(0), "0 штук")
}

func Test_NegativeOnePiece(t *testing.T) {
	assert.Equal(t, "штука", Format(-1), "-1 штука")
}

func Test_NegativeOneDigitShtuki(t *testing.T) {
	assert.Equal(t, "штуки", Format(-2), "-2 штуки")
	assert.Equal(t, "штуки", Format(-3), "-3 штуки")
	assert.Equal(t, "штуки", Format(-4), "-4 штуки")
}

func Test_NegativeOneDigitShtuk(t *testing.T) {
	assert.Equal(t, "штук", Format(-5), "-5 штук")
	assert.Equal(t, "штук", Format(-6), "-6 штук")
	assert.Equal(t, "штук", Format(-7), "-7 штук")
	assert.Equal(t, "штук", Format(-8), "-8 штук")
	assert.Equal(t, "штук", Format(-9), "-9 штук")
}

func Test_NegativeTeens(t *testing.T) {
	assert.Equal(t, "штук", Format(-11), "-11 штук")
	assert.Equal(t, "штук", Format(-12), "-12 штук")
	assert.Equal(t, "штук", Format(-13), "-13 штук")
	assert.Equal(t, "штук", Format(-14), "-14 штук")
	assert.Equal(t, "штук", Format(-15), "-15 штук")
	assert.Equal(t, "штук", Format(-16), "-16 штук")
	assert.Equal(t, "штук", Format(-17), "-17 штук")
	assert.Equal(t, "штук", Format(-18), "-18 штук")
	assert.Equal(t, "штук", Format(-19), "-19 штук")
}

func Test_NegativeLowerThan100(t *testing.T) {
	assert.Equal(t, "штук", Format(-27), "-27 штук")
	assert.Equal(t, "штук", Format(-45), "-45 штук")
	assert.Equal(t, "штука", Format(-61), "-61 штука")
	assert.Equal(t, "штуки", Format(-32), "-32 штуки")
	assert.Equal(t, "штуки", Format(-93), "-93 штуки")
	assert.Equal(t, "штуки", Format(-84), "-84 штуки")
}

func Test_NegativeTeensAfter100(t *testing.T) {
	assert.Equal(t, "штук", Format(-111), "-111 штук")
	assert.Equal(t, "штук", Format(-112), "-112 штук")
	assert.Equal(t, "штук", Format(-113), "-113 штук")
	assert.Equal(t, "штук", Format(-114), "-114 штук")
	assert.Equal(t, "штук", Format(-115), "-115 штук")
	assert.Equal(t, "штук", Format(-116), "-116 штук")
	assert.Equal(t, "штук", Format(-117), "-117 штук")
	assert.Equal(t, "штук", Format(-118), "-118 штук")
	assert.Equal(t, "штук", Format(-119), "-119 штук")
}

func Test_NegativeTeensAfter1000(t *testing.T) {
	assert.Equal(t, "штук", Format(-1011), "-1011 штук")
	assert.Equal(t, "штук", Format(-1012), "-1012 штук")
	assert.Equal(t, "штук", Format(-1013), "-1013 штук")
	assert.Equal(t, "штук", Format(-1014), "-1014 штук")
	assert.Equal(t, "штук", Format(-1015), "-1015 штук")
	assert.Equal(t, "штук", Format(-1016), "-1016 штук")
	assert.Equal(t, "штук", Format(-1017), "-1017 штук")
	assert.Equal(t, "штук", Format(-1018), "-1018 штук")
	assert.Equal(t, "штук", Format(-1019), "-1019 штук")
}

func Test_NearInt32LowerLimit(t *testing.T) {
	assert.Equal(t, "штука", Format(-2147483641), "-2147483641 штука")
	assert.Equal(t, "штуки", Format(-2147483644), "-2147483644 штуки")
	assert.Equal(t, "штук", Format(-2147483647), "-2147483647 штук")
	assert.Equal(t, "штук", Format(-2147483611), "-2147483611 штук")
}
