package main

import (
	"fmt"
	"log"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Исходный массив:", a)
	var n int
	fmt.Print("Номер удаляемого элемента: ")
	fmt.Scan(&n)
	if n < 0 || n >= len(a) {
		log.Fatal("Incorrect input value:", n)
	}

	copy(a[n:], a[n+1:])
	a[len(a)-1] = 0
	a = a[:len(a)-1]
	fmt.Println("Модифицированный массив:", a)
}
