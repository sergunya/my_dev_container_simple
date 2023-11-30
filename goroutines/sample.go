package main

import (
	"fmt"
	"time"
)

func main() {
	go waitAndPrint("Hello")
	_ = 1
	go waitAndPrint("Hello 2")
	_ = 2
	go waitAndPrint("Hello 3")
	_ = 3
	printAndWaitBit("Bye")
}

func waitAndPrint(msg string) {
	time.Sleep(time.Second)
	fmt.Println(msg)
}

func printAndWaitBit(msg string) {
	fmt.Println(msg)
	time.Sleep(2 * time.Second)
}
