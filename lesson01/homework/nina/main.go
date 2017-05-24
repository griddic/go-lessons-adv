package main

import (
	"fmt"
)

func main() {
	fmt.Println(Format(1009))
}

func Format(n int) string {
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
