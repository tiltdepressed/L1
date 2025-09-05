package main

import (
	"fmt"
	"sync"
	"time"
)

func Sleep(sleepTime time.Duration) {
	timer := time.NewTimer(sleepTime)
	defer timer.Stop()

	fmt.Println("Горутина спит...")
	<-timer.C
	fmt.Println("Горутина проснулась!")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 5 {
			fmt.Println(i)
			Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 5 {
			fmt.Println(-i)
			Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
}
