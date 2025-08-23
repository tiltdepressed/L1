package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	var n int
	fmt.Print("Введине N: ")
	fmt.Fscan(os.Stdin, &n)

	ch := make(chan int)

	timer := time.NewTimer(time.Duration(n) * time.Second)
	defer timer.Stop()

	go func() {
		for val := range ch {
			fmt.Println("Получено: ", val)
		}
	}()

	for {
		select {
		case <-timer.C:
			close(ch)
			return
		default:
			i := rand.IntN(100)
			ch <- i
			fmt.Println("Отправлено: ", i)
			time.Sleep(time.Second)
		}
	}
}
