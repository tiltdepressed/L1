package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mux sync.Mutex
	val int
}

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var counter Counter
	counter.val = 0

	const numGoroutines = 10
	wg.Add(numGoroutines)

	for range numGoroutines {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					counter.mux.Lock()
					counter.val += 1
					// fmt.Println("Текущее значение счетчика:", counter.val)
					counter.mux.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println("Значение счетвика:", counter.val)
}
