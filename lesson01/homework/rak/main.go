package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	fmt.Println(Format(num))
}

func Format(n int) string {
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
