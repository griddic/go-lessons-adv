package main

import (
	"fmt"
	"math"
)

func main() {
}

func FormatSize(s uint64) string {
	if s < 1000 {
		return fmt.Sprint(s) + "b"
	} else if s < 10000 {
		return fmt.Sprint(math.Floor(float64(s)/100.0)/10.0) + "Kb"
	} else {
		return fmt.Sprint(s/1000) + "Kb"
	}
}
