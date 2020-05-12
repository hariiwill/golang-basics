package main

import "fmt"

func main() {

	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}

	l := len(arr) + 1

	su := l * (l + 1) / 2

	fmt.Println(su)
	var sum = 0

	for i := 0; i < len(arr); i++ {

		sum = sum + arr[i]

	}

	fmt.Println(sum)

	res := su - sum

	fmt.Printf("Missing element is", res)
}
