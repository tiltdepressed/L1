package main

import (
	"fmt"
)

func main() {
	set := make(map[string]bool)
	result := []string{}

	var n int
	fmt.Print("Введите количество строк: ")
	if _, err := fmt.Scan(&n); err != nil || n < 0 {
		fmt.Println("Ошибка ввода количества элементов: ", n)
		return
	}

	var val string
	fmt.Print("Введите строки: ")
	for range n {
		fmt.Scan(&val)
		if !set[val] {
			set[val] = true
			result = append(result, val)
		}
	}
	fmt.Println(result)
}
