package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []card

func newDeck() deck {

	//declare some arrays
	cards := deck{}
	color := [4]string{"red", "blue", "yellow", "green"}
	faces := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	//Add Wild cards
	for i := 0; i < 5; i++ {
		var c card = createType1Card("Wild")
		cards = append(cards, c)
	}
	//Add WildDrawFour cards
	for i := 0; i < 5; i++ {
		var c card = createType1Card("WildDrawFour")
		cards = append(cards, c)
	}
	//Add type2 cards
	type2 := [4]string{"Skip", "Reverse", "DrawTwo"}
	for i := 0; i < 3; i++ { //skip reverse or drawtwo
		for j := 0; j < 4; j++ { //color
			for k := 0; k < 2; k++ { //two of each
				var c card = createType2Card(type2[i], color[j])
				cards = append(cards, c)
			}
		}
	}

	//Add regular cards (except 0)
	for i := 0; i < 4; i++ { //color
		for j := 1; j <= 9; j++ { //value
			var c card = createType3Card(faces[j], color[i])
			cards = append(cards, c)
			cards = append(cards, c)
		}
	}

	// Add zero
	for i := 0; i < 4; i++ {
		var c card = createType3Card("0", color[i])
		cards = append(cards, c)
	}

	return cards
}

func shuffle(deck deck) deck { //returns the shuffled deck
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func getTop(deck deck) card { //gets the top card from the deck
	// var c card = deck[len(deck)-1]
	// deck = deck[:len(deck)-1]
	// return c
	return deck[len(deck)-1]
}

func length(deck deck) int { //returns the number of cards in the deck
	return len(deck)
}

func removeTopCard(deck deck) deck { //removes the top card and returns the deck
	deck = deck[:len(deck)-1]
	return deck
}

func getTopCardAndReturnDeck(deck deck) (card, deck) { //removes the top card, returns the card and deck
	var c card = getTop(deck)
	deck = removeTopCard(deck)
	return c, deck
}

func dealCards(n int, d deck) (deck, deck) { //deals n cards into another deck and returns both the dealt cards as well as the original deck

	/*
		Returns:
		(dealt hand, original deck with removed cards from the top)
	*/
	tempHand := deck{}
	for i := 0; i < n; i++ {
		tempHand = append(tempHand, getTop(d))
		d = removeTopCard(d)
	}
	return tempHand, d
}

func drawTwo(d deck) (deck, deck) { //returns two cards (first deck) and the original deck with  two cards removed from the top
	return dealCards(2, d)
}

func drawfour(d deck) (deck, deck) { //returns four cards (first deck) and the original deck with  two cards removed from the top
	return dealCards(4, d)
}

func addCardToDeck(c card, d deck) deck { //returns a deck (the centre pile of cards where the players are placing their cards) after a player plays their card
	d = append(d, c)
	return d
}

func removeCardFromDeck(index int, d deck) deck {
	d[index] = d[len(d)-1]
	return d[:len(d)-1]
}

func printPlayerHand(p player) {
	for i, c := range p.hand {
		fmt.Printf("%v.", i)
		printCard(c)
		fmt.Printf("\n")
	}
}

func printDeck(d deck) {
	for i, c := range d {
		fmt.Printf("%v.", i)
		printCard(c)
		fmt.Printf("\n")
	}
}
