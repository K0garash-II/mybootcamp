package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func check(id int) (string, error) {
	delay := time.Duration(100+rand.Intn(900)) * time.Millisecond
	time.Sleep(delay)
	if id%7 == 0 {
		return "", fmt.Errorf("service %d: down", id)
	}
	return fmt.Sprintf("service %d: ok (%v)", id, delay), nil
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 5)
	result := make([]any, 20)
	ids := [20]int{}

	for i := 0; i < 20; i++ {
		ids[i] = i + 1
	}

	wg.Add(20)

	for i := 0; i < 20; i++ {
		ch <- struct{}{}
		go func(id int) {
			defer wg.Done()
			defer func() {<-ch}()
			res, err := check(id)
			if err != nil {
				result[id-1] = err
			} else {
				result[id-1] = res
			}
		}(ids[i])
	}

	wg.Wait()

	for _, v := range result {
		fmt.Println(v)
	}

	var calcNormal int
	var calcErr int

	for _, v := range result {
		switch v.(type) {
		case string:
			calcNormal += 1
		case error:
			calcErr += 1
		}
	}

	fmt.Printf("Normal services: %v\n", calcNormal)
	fmt.Printf("Down services: %v\n", calcErr)
}
