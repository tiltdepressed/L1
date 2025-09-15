package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/unicode/norm"
)

func main() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		input = norm.NFC.String(input)
		r := []rune(input)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		fmt.Println("Развернутая строка:", r)
	}
}
