package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("this is inside thread")
		wg.Done()
	}(&wg)
	fmt.Println("This is main function")
	wg.Wait()
}
