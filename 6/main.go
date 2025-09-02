package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 1. Выход по условию
	/*wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина завершается, когда заканчивается ее функция или выполняется условие")
	}()*/

	// 2. Канал уведомления
	/*sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for {
			select {
			case <-sigCh:
				fmt.Println("Горутина завершилась через канал уведомения")
				return
			default:
				count += 1
			}
		}
	}()
	<-sigCh
	close(sigCh)*/

	// 3. Через контекст
	/*ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершилась через контекст")
				return
			default:
				count += 1
			}
		}
	}()
	<-sigCh
	close(sigCh)
	cancel()*/

	// 4. RuntimeGoExit (прерывает горутину, но не выходит из нее)
	/*go func() {
		fmt.Println("Здесь какая-то логика")
		runtime.Goexit()
		fmt.Println("Этого сообщения мы не увидим")
	}()
	*/

	// 5. Panic
	/*defer func() {
		fmt.Println("main defer")
	}()
	wg.Add(2)
	go func() {
		defer func() {
			fmt.Println("first defer")
			wg.Done()
		}()

		go func() {
			defer func() {
				fmt.Println("second defer")
				wg.Done()

				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			panic("Something went wrong(((")
		}()

		time.Sleep(time.Second)
	}()*/

	// 6. Закрытие канала
	/*teshCh := make(chan string)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("6 завершилась закрытием канала")
		}()
		for val := range teshCh {
			fmt.Println(val)
		}
	}()
	teshCh <- "testValue"
	close(teshCh)
	*/
	wg.Wait()

	// 7. Завершение main функции
	// Когда main функция сама завершается, все остальные принудительно останавливаются
}
