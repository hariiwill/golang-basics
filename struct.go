package main

import "fmt"

type simpleInterface interface {
	print()
}

type SimpleStruct struct {
	Name   string
	Age    int
	Gender bool
}

func print(s []SimpleStruct) {

	for _, v := range s {

		fmt.Println(v)
	}
}

func main() {

	var stringaccess SimpleStruct
	var stringaccess1 SimpleStruct
	stringaccess.Name = "Subhi"
	stringaccess.Age = 26
	stringaccess.Gender = true
	//stringaccess.print()

	stringaccess1.Name = "vihaan"
	stringaccess1.Age = 2
	stringaccess1.Gender = true
	//stringaccess1.print()

	//fmt.Println(stringaccess)

	var s []SimpleStruct
	s = append(s, stringaccess, stringaccess1)
	print(s)

}
