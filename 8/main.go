package main

import (
	"fmt"
)

func main() {
	var input int64
	var number uint
	var value bool
	fmt.Println("Введите число, номер бита и его значение в виде 'x y z' (отсчет ведется с нуля):")
	if _, err := fmt.Scan(&input, &number, &value); err != nil {
		fmt.Println(err)
		return
	}

	if number > 63 {
		fmt.Println("Ошибка: номер бита должен быть от 0 до 63")
		return
	}

	if value {
		input = input | (1 << number)
	} else {
		input = input &^ (1 << number)
	}

	fmt.Println("Новое значение:", input)
}
