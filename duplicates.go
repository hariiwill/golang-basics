package main

import "fmt"

func removeDuplicates(arr []int, i int) []int {

	return append(arr[:i], arr[i+1:]...)
}

func main() {

	arr := []int{1, 2, 3, 4, 4, 5, 6, 7, 2, 1}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		for j := i + 1; j < len(arr); j++ {

			if arr[i] == arr[j] {

				arr = removeDuplicates(arr, j)

			}

		}

	}

	fmt.Println(arr)
}
