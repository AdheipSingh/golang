package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var globalvar int

func main() {
	wg.Add(2)
	go inc("one:")
	go inc("two")
	wg.Wait()

	fmt.Println("Globalvar: ", globalvar)

}

func inc(s string) {
	for i := 0; i < 10; i++ {
		a := globalvar
		a++
		globalvar = a
		fmt.Println(s, i, "globalvar", globalvar)

	}
	wg.Done()
}
