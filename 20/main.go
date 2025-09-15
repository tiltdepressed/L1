package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/unicode/norm"
)

func reverseRunes(r []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func main() {
	fmt.Print("Введите предложение: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		sentence := scanner.Text()
		sentence = norm.NFC.String(sentence)
		r := []rune(sentence)
		reverseRunes(r, 0, len(r)-1)

		wordStart := 0
		for i := 0; i <= len(r); i++ {
			if i == len(r) || r[i] == ' ' {
				reverseRunes(r, wordStart, i-1)
				wordStart = i + 1
			}
		}
		fmt.Println("Развернутое предложение:", string(r))
	}
}
