package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{100}

	for i := 0; i < 5; i++ {
		arr = append(arr, rand.Intn(50))
	}

	fmt.Println("Initial", arr)

	bubbleSort(arr)

	fmt.Println("Final", arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {

				jack := arr[j]

				arr[j], arr[j+1] = arr[j+1], jack
			}

			fmt.Println(arr)
		}


	}
}