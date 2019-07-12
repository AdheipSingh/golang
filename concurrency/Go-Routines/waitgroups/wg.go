package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go one()
	go two()
	wg.Wait()
}

func one() {
	for i := 0; i < 100; i++ {
		fmt.Println("One", i)
	}
	wg.Done()
}

func two() {
	for i := 0; i < 100; i++ {
		fmt.Println("Two", i)
	}
	wg.Done()
}
