package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, val := range arr {
		groups[int(val)/10*10] = append(groups[int(val)/10*10], val)
	}

	fmt.Println(groups)
}
