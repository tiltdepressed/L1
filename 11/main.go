package main

import (
	"fmt"
	"log"
)

func makeSet(num int) map[int]bool {
	var n, val int
	fmt.Printf("Введите кол-то элеметнов множества %d: ", num)
	if _, err := fmt.Scan(&n); err != nil || n < 0 {
		log.Fatal("Ошибка ввода количества элементов: ", n)
	}

	set := make(map[int]bool)
	fmt.Printf("Введите %d элементов множества (int): ", n)
	for range n {
		if _, err := fmt.Scan(&val); err != nil {
			log.Fatal("Ошибка ввода числа: ", val)
		}
		set[val] = true
	}
	return set
}

func main() {
	set1 := makeSet(1)
	set2 := makeSet(2)
	res := []int{}

	for i := range set1 {
		if set2[i] {
			res = append(res, i)
		}
	}
	fmt.Println("Пересечение множеств:", res)
}
