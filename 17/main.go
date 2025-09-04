package main

import (
	"fmt"
	"math/rand"
)

func binSearch(arr []int, val, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == val {
		return mid
	} else if arr[mid] < val {
		return binSearch(arr, val, mid+1, right)
	} else {
		return binSearch(arr, val, left, mid-1)
	}
}

func main() {
	arr := make([]int, 10)
	for i := range len(arr) {
		arr[i] = rand.Intn(20)
	}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Массив:", arr)
	val := rand.Intn(20)
	fmt.Println("Искомое число:", val)
	index := binSearch(arr, val, 0, len(arr)-1)
	fmt.Println("Индекс искомого элемента:", index)
}

func quickSort(arr []int, low, high int) {
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
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}
	if j > low {
		quickSort(arr, low, j)
	}
	if i < high {
		quickSort(arr, i, high)
	}
}
