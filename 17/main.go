package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.IntN(20)
	}

	slices.Sort(arr)

	fmt.Println("Отсортированный массив:", arr)

	val := rand.IntN(20)
	fmt.Println("Искомое число:", val)

	index := binarySearch(arr, val)

	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу: %d\n", val, index)
	} else {
		fmt.Printf("Элемент %d в массиве не найден.\n", val)
	}
}
