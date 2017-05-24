package main

import (
	"fmt"
	"math"
)

// func main() {
//
// 	for i := -30; i <= 30; i++ {
// 		fmt.Println(i, Format(i))
// 	}
// }

func main() {
	fmt.Println(Format(1))
}

func Format(n int) string {
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
