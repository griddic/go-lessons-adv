package main

import (
	"fmt"
	"math"
)

func main() {

}

// Round стащил из интернета и немного поправил по себя.
func Round(val float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// FormatSize ... uint64 больше чем на Pb не хватает...
func FormatSize(s uint64) string {
	input := float64(s)
	index := []string{"Kb", "Mb", "Gb", "Tb", "Pb", "Eb", "Zb", "Yb"}
	prefixIndex := 0

	if input < 1000 {
		return fmt.Sprint(s) + "b"
	}

	// Выше 999500 выводить следующий префикс (будет округлено вверх до единицы следующего префикса)
	for input >= 999500 {
		input = input / 1000
		prefixIndex++
	}

	input = Round(input, 1)

	if input < 10000 {
		return fmt.Sprintf("%.2f", input/1000) + index[prefixIndex]
	} else if input < 100000 {
		return fmt.Sprintf("%.1f", input/1000) + index[prefixIndex]
	} else if input < 1000000 {
		return fmt.Sprintf("%.0f", input/1000) + index[prefixIndex]
	}

	return fmt.Sprint("Something went wrong. Number provided=", s)

}
