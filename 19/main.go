package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		r := []rune(input)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		fmt.Println("Развернутая строка:", input)
	}
}
