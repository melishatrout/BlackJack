package main

import (
	"blackjack"
	"fmt"
)

func main() {
	ga := blackjack.New()
	winnings := ga.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}