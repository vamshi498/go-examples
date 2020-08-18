package main

import (
	"fmt"
	"sync"
)

//buffered channel
var sem = make(chan int, 2)
var wg sync.WaitGroup

func main() {
	fmt.Println("hello")

	wg.Add(2)

	go producer()
	go consumer()

	wg.Wait()

}

func producer() {
	defer wg.Done()
	defer close(sem)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range nums {
		sem <- v

	}
}

func consumer() {
	defer wg.Done()
	for c := range sem {
		fmt.Printf("value is %d\n", c)
	}

}
