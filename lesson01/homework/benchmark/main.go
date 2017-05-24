package main

import (
	"math"
	"strconv"
)

// Fmorte
func FormatFmorte(n int) string {
	switch Abs(n) % 100 {
	case 11, 12, 13, 14:
		return "штук"
	}
	switch Abs(n) % 10 {
	case 0, 5, 6, 7, 8, 9:
		return "штук"
	case 2, 3, 4:
		return "штуки"
	}
	return "штука"
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}
	return x
}

// Grigory
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func FormatGrigory(n int) string {
	abs_n := abs(n)
	last_two_digits := abs_n % 100
	if (last_two_digits > 4) && (last_two_digits < 21) {
		return "штук"
	}
	kvantificators := []string{
		"штук",  // 0
		"штука", // 1
		"штуки", // 2
		"штуки", // 3
		"штуки", // 4
		"штук",  // 5
		"штук",  // 6
		"штук",  // 7
		"штук",  // 8
		"штук",  // 9
	}

	var last_digit int = abs_n % 10
	return kvantificators[last_digit]
}

// Kislov

func FormatKislov(n int) string {
	b := strconv.Itoa(n)
	if len(b) > 1 {
		b = b[len(b)-2:]
	}
	c, _ := strconv.Atoi(b)
	if c >= 11 && c >= 11 && c <= 14 {
		return "штук"
	} else {
		b := strconv.Itoa(n)
		b = b[len(b)-1:]
		c, _ := strconv.Atoi(b)
		switch c {
		case 1:
			return "штука"
		case 2, 3, 4:
			return "штуки"
		case 5, 6, 7, 8, 9, 0:
			return "штук"
		}
	}
	return "штук"
}

// Kostechuk
func FormatKostechuk(n int) string {
	n = int(math.Abs(float64(n % 100)))
	if n >= 5 && n <= 19 {
		return "штук"
	}

	n = int(math.Abs(float64(n % 10)))
	if n == 1 {
		return "штука"
	} else if n >= 2 && n <= 4 {
		return "штуки"
	} else {
		return "штук"
	}
}

// Rak
func FormatRak(n int) string {
	if n < 0 {
		n = n * -1
	}
	rest10 := n % 10
	rest100 := n % 100
	var unit string = "штуксель"

	switch {
	case (5 <= rest10 && rest10 <= 9) || rest10 == 0 || (11 <= rest100 && rest100 <= 14):
		unit = "штук"
	case rest10 == 1:
		unit = "штука"
	case 2 <= rest10 && rest10 <= 4:
		unit = "штуки"
	}

	return unit
}

// Nikolay
func FormatNikolay(n int) string {
	last_digits := int(math.Abs(float64(n % 100)))
	if (11 <= last_digits) && (last_digits <= 19) {
		return "штук"
	}
	last_digits %= 10
	switch last_digits {
	case 1:
		return "штука"
	case 2:
		return "штуки"
	case 3:
		return "штуки"
	case 4:
		return "штуки"
	default:
		return "штук"
	}
}

// Nina
func FormatNina(n int) string {
	if n < 0 {
		n = -n
	}
	if n%10 >= 5 {
		return "штук"
	} else if n%100 >= 05 && n%100 <= 21 {
		return "штук"
	} else if n%100 == 01 {
		return "штука"
	} else if n%10 == 1 {
		return "штука"
	} else if n%10 == 0 || n%100 == 00 {
		return "штук"
	}
	return "штуки"
}
