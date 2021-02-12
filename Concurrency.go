// Created as a simple program to understand goroutines
// Repeatedly prints what you type without waiting to scan

package main

import (
	"fmt"
)


var message string = "test"

func main() {
	go multiplyBigNumbers()
	for {
		fmt.Println(message)
		if message == "exit"{
			break
		}
	}
}

func multiplyBigNumbers() {
	fmt.Scanln(&message)
	go multiplyBigNumbers()
}
