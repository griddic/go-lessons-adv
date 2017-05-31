package main

import (
	"fmt"
	"math"
)

func main() {

}

func FormatSize(s uint64) string {
	if s < 999 {
		return fmt.Sprint(s) + "b"
	} else if s < 999950 {
		return fmt.Sprint(math.Floor(float64(s)/1000.0*10.0+0.5)/10.0) + "Kb"
	} else if s < 999950000 {
		return fmt.Sprint(math.Floor(float64(s)/1000000.0*10.0+0.5)/10.0) + "Mb"
	} else {
		return fmt.Sprint(math.Floor(float64(s)/1000000000.0*10.0+0.5)/10.0) + "Gb"
	}
}
