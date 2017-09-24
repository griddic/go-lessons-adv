package main

import "fmt"

//Nameable dfsdf
type Nameable interface {
	Name() string
}

type User struct {
	id   int64
	name string
}

func (u *User) Name() string {
	return u.name
}

type City struct {
	country string
	name    string
}

func (u *City) Name() string {
	return u.name
}

func main() {
	u := User{name: "Alex"}
	c := City{name: "Moscow"}

	processObjectName(&u)
	processObjectName(&c)
}

func processObjectName(obj Nameable) {
	fmt.Println("The name is:", obj.Name())
}
