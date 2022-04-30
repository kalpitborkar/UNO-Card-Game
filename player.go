package main

type player struct {
	name string
	hand deck
}

type players []player

/*
Need to set a global variable "current player" in logic.go and increment it to move to the next player
currentPlayer = (currentPlayer+1)%len(players)

Idea:
Can have a string array consisting of the names of the players
*/
