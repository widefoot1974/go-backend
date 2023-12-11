package main

import "fmt"

func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	go squareIt(inputChannel, outputChannel)
	for i := 0; i < 10; i++ {
		inputChannel <- i
		j := outputChannel
		fmt.Println(j)
	}
	// for i := range outputChannel {
	// 	fmt.Println(i)
	// }
}

func squareIt(inputChan, outputChan chan int) {
	for x := range inputChan {
		outputChan <- x * x
	}
}
