package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	standardBoard := MakeStandardBoard()
	fmt.Println(standardBoard.GetMoves())
	fmt.Println(standardBoard)
}
