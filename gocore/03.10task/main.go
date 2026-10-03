package main

import (
	"fmt"
	"sync"
)

func main() {
	
}

func Async() {
	var wg sync.WaitGroup
	for i := 0; i <= 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Малайка балалайка гомик, поздравляю")
		}()
	}
}

func Conv() {
	for i := 0; i <= 100000; i++ {
		fmt.Println("Малайка балалайка гомик, поздравляю")
	}
}