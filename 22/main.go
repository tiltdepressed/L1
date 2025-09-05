package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("23456789012345678901234567890", 10)
	b, _ := new(big.Int).SetString("98765432109876543210987654321", 10)

	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("A + B =", new(big.Int).Add(a, b))
	fmt.Println("A - B =", new(big.Int).Sub(a, b))
	fmt.Println("A / B =", new(big.Int).Div(a, b))
	fmt.Println("A * B =", new(big.Int).Mul(a, b))
}
