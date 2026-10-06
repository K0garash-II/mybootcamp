package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	x := rand.Float64()
	fmt.Printf("%s processing order 1\n", name)
	time.Sleep(time.Duration(x * float64(time.Second)))
	fmt.Printf("%s processing order 2\n", name)
	time.Sleep(time.Duration(x * float64(time.Second)))
	fmt.Printf("%s processing order 3\n", name)
	time.Sleep(time.Duration(x * float64(time.Second)))

	fmt.Printf("%s finished\n", name)
}

func calculate(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := num * num
	ch <- res
}

func main() {
	var wg sync.WaitGroup

	nums := []int{10, 20, 30}

	ch := make(chan int, 3)

	wg.Add(3)

	go calculate(nums[0], ch, &wg)
	go calculate(nums[1], ch, &wg)
	go calculate(nums[2], ch, &wg)

	value := <-ch
	fmt.Println(value)
	value = <- ch
	fmt.Println(value)
	value = <- ch
	fmt.Println(value)

	wg.Wait()
}
