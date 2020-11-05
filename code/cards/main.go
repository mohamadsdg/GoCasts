package main

import "fmt"

func main(){
	// var card string = "Ace of Spedes"
	// card := "Ace of Spedes" // declare and assing
	//card = "ddfdf" // reassign
	// card := newCard() 
	cards := []string{newCard(),"Ace of Diamonds"}
	cards = append(cards,"Six of Spedes") //return new slice and reassign to prev variable
	for i,card := range cards{
		fmt.Println(i,card)
	}
}

func newCard() string{
	return "five of Diamond"
}