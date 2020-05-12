package main

import "fmt"

func changeString(s ...string) {

	s[1] = "Go"
	s = append(s, "playground")
	fmt.Println(s)

}

func main() {

	string1 := []string{"Hello", "world"}
	changeString(string1...)
	fmt.Println("Main", string1)
}
