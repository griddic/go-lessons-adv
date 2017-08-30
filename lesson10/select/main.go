package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan string, 5)

	go func() {
		for {
			c1 <- rand.Intn(20)
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		}
	}()

	go func() {
		for {
			c2 <- "Whazuup"
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		select {
		case num := <-c1:
			fmt.Println("Num:", num)
		case <-time.After(time.Second * 2):
			fmt.Println("Timeout!")
		}
	}
}
