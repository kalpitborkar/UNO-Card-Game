package main

import "fmt"

var numberOfPlayers int
var playerNames []string

// Create an object of the struct 'players' - namely 'allPlayers'
var allPlayers = players{}
var allPlayerHands []deck



allPlayerHands = createPlayerHandsForAllPlayers(numberOfPlayers, 5, mainDeck) 
allPlayers = createAllPlayersList(allPlayerNames, allPlayerHands)

func welcomeToUNO() {
	fmt.Printf("Welcome to UNO!\n")
	fmt.Println("Enter '1' to start the game")
	fmt.Println("Enter '2' to see the instructions")
	fmt.Println("Enter '3' to exit the game")
}

func getPlayerInfo() {
	fmt.Println("Enter the number of players.")
	fmt.Scanln(&numberOfPlayers)

	for i := 0; i < numberOfPlayers; i++ {
		fmt.Printf("Enter name of player %v.\n", i)
		var tempName string
		fmt.Scanln(&tempName)
		playerNames = append(playerNames, tempName)
	}
}

func printPlayerInfo() {
	fmt.Printf("Number of players: %v\n", numberOfPlayers)
	for i, name := range playerNames {
		fmt.Printf("Player %v: %v\n", i, name)
	}
}

func createPlayerObject(name string, playerHand deck) player { //takes a name and a hand, creates and returns a player object
	var tempPlayer player
	tempPlayer.hand = playerHand
	tempPlayer.name = name
	return tempPlayer
}

func createPlayerHandsForAllPlayers(n int, m int, d deck) []deck {
	/*
		n - number of players
		m - number of cards in each hand
		d - main deck containing all the cards at the beginning
		returns a slice of deck containing hands for each player
	*/
	var deckAllPlayerHands []deck
	for i := 0; i < n; i++ {
		tempHand := deck{}
		tempHand, d = dealCards(m, d)
		deckAllPlayerHands = append(deckAllPlayerHands, tempHand)
	}
	return deckAllPlayerHands
}

func createAllPlayersList(allPlayerNames []string, allPlayerHands []deck) players {
	//Returns a list of all player objects
	var allPlayersList = players{}
	for i := 0; i < numberOfPlayers; i++ {
		var playerObj player
		playerObj.name = allPlayerNames[i]
		playerObj.hand = allPlayerHands[i]
		allPlayersList = append(allPlayersList, playerObj)
	}
	return allPlayersList
}
