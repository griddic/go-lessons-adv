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

func (u *User) String() string {
	return fmt.Sprintf("User with the name %s and id %d", u.name, u.id)
}

type City struct {
	country string
	name    string
}

func (u *City) Name() string {
	return u.name
}

func main() {
	u := User{id: 234, name: "Alex"}
	c := City{name: "Moscow", country: "Russia"}

	fmt.Println(&u)
	fmt.Println(c)
}
