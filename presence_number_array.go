package main

import "fmt"

func presenceNumber(arr []int, s int) (int, int) {

	var sum, num1, num2 int

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {
			sum = arr[i] + arr[j]

			if s == sum {
				num1 = arr[i]
				num2 = arr[j]
			}

		}
	}

	return num1, num2

}

func main() {

	arr := []int{5, 6, 7, 3, 2, 0, 1, 5}

	n1, n2 := presenceNumber(arr, 13)

	fmt.Println(n1, n2)

}
