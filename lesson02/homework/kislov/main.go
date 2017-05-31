package main

import (
	"fmt"
	"math"
)

func main() {

}

func FormatSize(s uint64) string {
	if s < 1024 {
		return fmt.Sprint(s) + "b"
	} else if s < 100000 {
		return fmt.Sprint(math.Floor(float64(s)/100.0)/10.0) + "KB"
	} else if s < 1048576{
		return fmt.Sprint(s/1024) + "KB"
	} else if s == 1048576 {
		return fmt.Sprint(s/1048576) + "MB"
	} else if s < 104857600 {
		return fmt.Sprint(math.Floor(float64(s)/100000.0)/10.0) + "MB"
	} else if s < 1073741824 {
		return fmt.Sprint(s/1048576) + "MB"
	} else if s == 1073741824 {
		return fmt.Sprint(s/1073741824) + "GB"
	} else if s < 107374182400 {
		return fmt.Sprint(math.Floor(float64(s)/100000000.0)/10.0) + "GB"
	} else {
		return fmt.Sprint(s/1073741824) + "GB"
	}

}
