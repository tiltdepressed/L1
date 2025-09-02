package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	var n int
	fmt.Print("Введите N: ")
	fmt.Fscan(os.Stdin, &n)

	ch := make(chan int)
	for i := range n {
		go func(c chan int, num int) {
			for val := range ch {
				fmt.Println(num, ": ", val)
			}
		}(ch, i)
	}

	for {
		ch <- rand.IntN(100)
		time.Sleep(1000 * time.Millisecond)
	}
}
