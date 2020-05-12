package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "vihaansubhi"

	split := strings.Split(input, "")

	fmt.Println(split)

	var result string

	for _, v := range split {

		result = string(v) + result

		fmt.Println(result)

	}

}
