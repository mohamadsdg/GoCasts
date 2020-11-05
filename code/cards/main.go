package main

import "fmt"

func main(){
	// var card string = "Ace of Spedes"
	// card := "Ace of Spedes" // declare and assing
	//card = "ddfdf" // reassign
	card := newCard() // declare and assing
	fmt.Println(card)
}

func newCard() string{
	return "five of Diamond"
}