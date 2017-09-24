package main

import (
	"fmt"
	"sort"
	"strings"
)

type User struct {
	Id   int
	Name string
}

type ById []User

func (u ById) Len() int {
	return len(u)
}

func (u ById) Less(i, j int) bool {
	return u[i].Id < u[j].Id
}

func (u ById) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type ByName []User

func (u ByName) Len() int {
	return len(u)
}

func (u ByName) Less(i, j int) bool {
	return strings.Compare(u[i].Name, u[j].Name) < 0
}

func (u ByName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{2, "John"},
		{34, "Peter"},
		{1, "Kevin"},
		{23, "Nancy"},
	}

	fmt.Println(users)
	sort.Sort(ById(users))
	fmt.Println(users)
	sort.Stable(ByName(users))
	fmt.Println(users)
}
