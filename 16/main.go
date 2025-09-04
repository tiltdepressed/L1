package main

import (
	"fmt"
	"math/rand"
)

func quickSortRecursive(arr []int, low, high int) {
	i := low
	j := high
	pivot := arr[(i+j)/2]

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if j > low {
		quickSortRecursive(arr, low, j)
	}
	if i < high {
		quickSortRecursive(arr, i, high)
	}
}

func quickSort(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	quickSortRecursive(newArr, 0, len(newArr)-1)
	return newArr
}

func main() {
	arr := make([]int, 20)
	for i := range len(arr) {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	arrSorted := quickSort(arr)
	fmt.Println(arrSorted)
}
