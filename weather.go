package main

import (
	"fmt"
	"sync"
)

var readChannel = make(chan int)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go readData(readChannel, &wg)
	go printData(&wg)

	wg.Wait()
}

func printData(wg *sync.WaitGroup) {

	for element := range readChannel {
		fmt.Println("The element is", element)
	}
	wg.Done()

}

func readData(readChannel chan int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		readChannel <- i
	}

	close(readChannel)

}
