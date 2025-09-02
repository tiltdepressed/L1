package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))

	for _, val := range arr {
		go func(v int) {
			ch <- v * v
		}(val)
	}

	res := make([]int, len(arr))
	for i := range res {
		res[i] = <-ch
	}

	fmt.Println(res)
}
