package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	for {
		fmt.Print("Введите целое число: ")
		_, err := fmt.Scanln(&n)

		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Ошибка:", err.Error())
		} else {
			break
		}
	}
	fmt.Println(Format(n))
}

func Format(n int) string {
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
