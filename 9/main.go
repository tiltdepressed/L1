package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func makeArr(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func main() {
	var wg sync.WaitGroup
	first := make(chan int)
	second := make(chan int)

	arr := makeArr(10)
	fmt.Println(arr)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, val := range arr {
			first <- arr[val]
		}
		close(first)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range first {
			second <- val * 2
		}
		close(second)
	}()

	for res := range second {
		fmt.Println(res)
	}

	wg.Wait()
}
