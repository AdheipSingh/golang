package main

import (
	"fmt"
)

func main() {
	go one()
	go two()
}

func one() {
	for i := 0; i < 100; i++ {
		fmt.Println("One", i)
	}
}

func two() {
	for i := 0; i < 100; i++ {
		fmt.Println("Two", i)
	}
}
