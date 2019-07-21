package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
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
		mutex.Lock()
		a := globalvar
		a++
		globalvar = a
		fmt.Println(s, i, "globalvar", globalvar)
		mutex.Unlock()
	}
	wg.Done()
}
