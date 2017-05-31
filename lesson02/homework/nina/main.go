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
	} else if s < 1000000 {
		return fmt.Sprint(math.Floor(float64(s)/1000.0*10.0+0.5)/10) + "Kb"
	} else if s < 1000000000 {
		return fmt.Sprint(math.Floor(float64(s)/1000000.0*10.0+0.5)/10) + "Mb"
	} else if s < 100000000000 {
		return fmt.Sprint(math.Floor(float64(s)/1000000000.0*10.0+0.5)/10) + "Gb"
	} else {
		return "Keep calm and code on GO"
	}

}
