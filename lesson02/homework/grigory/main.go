package main

import (
	"fmt"
	"math"
)

var suffixes = []string{"B", "KB", "MB", "GB"}

func main() {}

func humanize(n uint64) string {
	length := uint(math.Floor(math.Log10(float64(n))))
	if length > 11 {
		length = 11
	}

	var divider uint64 = 1
	var point_location uint = 0
	if length > 2 {
		divider = uint64(math.Pow10(int(length) - 2))
		point_location = (2 * (length + 1)) % 3
	}

	suffix_index := length / 3
	to_print_float := float64(n/divider) / float64(math.Pow10(int(point_location)))

	return fmt.Sprint(to_print_float) + suffixes[suffix_index]
}
