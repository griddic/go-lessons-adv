package main

import (
    "fmt"
)

func main() {
    fmt.Println(Format(0))
    fmt.Println(Format(1))
    fmt.Println(Format(2))
    fmt.Println(Format(3))
    fmt.Println(Format(4))
    fmt.Println(Format(5))
    fmt.Println(Format(6))
    fmt.Println(Format(7))
    fmt.Println(Format(8))
    fmt.Println(Format(9))
    fmt.Println(Format(10))
    fmt.Println(Format(11))
    fmt.Println(Format(12))
    fmt.Println(Format(13))
    fmt.Println(Format(14))
    fmt.Println(Format(15))
    fmt.Println(Format(16))
    fmt.Println(Format(17))
    fmt.Println(Format(18))
    fmt.Println(Format(19))
    fmt.Println(Format(20))
    fmt.Println(Format(21))
    fmt.Println(Format(22))
    fmt.Println(Format(222225))
    fmt.Println(Format(-2))
    fmt.Println(Format(-11))
    fmt.Println(Format(-0))
    fmt.Println(Format(-1001))
}

func Format(n int) string {
    switch Abs(n) % 100 {
    case 11, 12, 13, 14:
        return "штук"
    }
    switch Abs(n) % 10 {
    case 0, 5, 6, 7, 8, 9:
        return "штук"
    case 2, 3, 4:
        return "штуки"
    }
    return "штука"
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }

    if x == 0 {
        return 0
    }
    return x
}
