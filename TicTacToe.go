package main

import (
	"fmt"
	"strconv"
)

var board [9]string
var turn string = "X"

//switches between X and O
func changeTurn() {
	if turn == "X" {
		turn = "O"
	} else {
		turn = "X"
	}
}
//checks whether the board is in a win state
func checkWin() bool {
	//checks rows and cols
	for i := 0; i < 3; i++ {
		if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] {
			return true
		}
		if board[i] == board[i+3] && board[i+3] == board[i+6] {
			return true
		}
	}
	if board[0] == board[4] && board[4] == board[8] {
		return true
	}
	if board[2] == board[4] && board[4] == board[6] {
		return true
	}
	return false
}
func checkTie(){
	
}
func main() {
	for i := range board {
		board[i] = strconv.Itoa(i)
	}
	fmt.Printf("It is %s's turn. Type a number from 0 to 8.\n", turn)
	for {
		fmt.Printf("%v\n%v\n%v\n", board[0:3], board[3:6], board[6:])
		var selection int
		n, err := fmt.Scanf("%d\r\n", &selection) // I have Windows so this may not work on a Linux machine because of the \r\n
		if selection >= 0 && selection <= 8 && err == nil && n == 1 && board[selection] != "X" && board[selection] != "O" {
			board[selection] = turn
			if checkWin() {
				fmt.Printf("%v\n%v\n%v\n", board[0:3], board[3:6], board[6:])
				fmt.Printf("%s wins!", turn)
				break
			}
			changeTurn()

			fmt.Printf("It is %s's turn. Type a number from 0 to 8.\n", turn)
		} else {
			fmt.Printf("Incorrect input. Type a number from 0 to 8.\n")
		}
	}

}
