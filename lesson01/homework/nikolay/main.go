package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Format(1))
}

func Format(n int) string {
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