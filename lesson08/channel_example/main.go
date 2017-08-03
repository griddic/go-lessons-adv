package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	input := make(chan int)
	output := make(chan int)

	go worker("first", input, output)
	go worker("second", input, output)
	go worker("third", input, output)

	numCount := 20
	go func() {
		for i := 0; i < numCount; i++ {
			input <- rand.Intn(20)
		}
		close(input)
	}()

	for i := 0; i < numCount; i++ {
		fmt.Println("out:", <-output)
	}
	time.Sleep(time.Second * 2)
}

func worker(name string, input chan (int), output chan (int)) {
	fmt.Println("started", name)
	for num := range input {
		fmt.Println(name, num)
		output <- num + 1
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
	fmt.Println("finished", name)
}
