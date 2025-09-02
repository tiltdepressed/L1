package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) String() string {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return fmt.Sprint(sm.m)
}

func randWord() string {
	return string([]byte{byte(rand.Intn(26) + 'a'), byte(rand.Intn(26) + 'a'), byte(rand.Intn(26) + 'a')})
}

func main() {
	var wg sync.WaitGroup

	sm := NewSafeMap()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			word := randWord()
			val := rand.Intn(100)
			sm.Set(word, val)
			fmt.Println("1 - ", word, ":", val)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			word := randWord()
			val := rand.Intn(100)
			sm.Set(word, val)
			fmt.Println("2 - ", word, ":", val)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println(sm.String())
}
