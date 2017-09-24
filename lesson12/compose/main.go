package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Coords struct {
	x float64
	y float64
}

type User struct {
	Coords
	Id   int
	Name string
}

func main() {
	u := User{}
	u.x = 3
	fmt.Println(u)
}
