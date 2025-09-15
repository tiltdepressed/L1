package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

func isUniqueString(input string) bool {
	input = strings.ToLower(input)
	symbolCount := make(map[rune]bool)

	for _, char := range input {
		if symbolCount[char] {
			return false
		}
		symbolCount[char] = true
	}
	return true
}

func main() {
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		input = norm.NFC.String(input)

		fmt.Println(isUniqueString(input))
	}
}
