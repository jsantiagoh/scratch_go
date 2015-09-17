package main

import (
	"fmt"
	"math/rand"
	"time"
)

func callA() int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
	return 40
}

func callB() int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))
	return 2
}

func question() int {
	fmt.Println("The question")

	c := make(chan int, 2)
	result := make(chan int)

	go func() {
		c <- callA()
	}()
	go func() {
		c <- callB()
	}()
	go func() {
		result <- (<-c) + (<-c)
	}()

	select {
	case total := <-result:
		return total
	case <-time.After(time.Second * 2):
		panic("no response after 2 seconds")
	}
}

func main() {
	fmt.Println(question())
}
