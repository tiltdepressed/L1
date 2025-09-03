package main

import "fmt"

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a := 5
	b := 13

	fmt.Printf("old a: %d, old b: %d\n", a, b)
	fmt.Println("SWAP!")
	swap(&a, &b)
	fmt.Printf("new a: %d, new b: %d\n", a, b)
}
