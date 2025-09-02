package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	sigCh := make(chan os.Signal, 1) // Канал считывает Ctrl+C
	signal.Notify(sigCh, os.Interrupt)

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for val := range ch {
				fmt.Println(num, ": ", val)
			}
			fmt.Println("Завершилась горутина ", num)
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigCh:
				close(ch) // Когда закрывается этот канал, завершаются все горутины-воркеры
				return    // Завершаем и текущую горутину
			default:
				ch <- rand.IntN(100)
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()
	<-sigCh
	close(sigCh) // Канал больше не ждет сигналов и начинается закрытие горутин
	wg.Wait()
}
