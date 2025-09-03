package main

import "fmt"

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	diapazoni := make(map[int][]float32)

	for _, val := range arr {
		diapazoni[int(val)/10*10] = append(diapazoni[int(val)/10*10], val)
	}

	fmt.Println(diapazoni)
}
