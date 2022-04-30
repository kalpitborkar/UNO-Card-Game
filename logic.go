package main

import "fmt"

func allowedPlayCard(playerCard card, topCentrePileCard card) bool {
	if playerCard.cardType == "WildDrawFour" {
		return true
	}
	if topCentrePileCard.cardType == "WildDrawFour" {
		if playerCard.cardType == "WildDrawFour" {
			return true
		} else {
			return false
		}
	}
	if topCentrePileCard.cardType == "DrawTwo" {

		if playerCard.cardType == "DrawTwo" || playerCard.cardType == "WildDrawFour" {
			return true
		} else {
			return false
		}
	}
	if topCentrePileCard.cardType == "Skip" || topCentrePileCard.cardType == "Reverse" {
		if playerCard.color == topCentrePileCard.color || playerCard.cardType == topCentrePileCard.cardType {
			return true
		} else {
			return false
		}
	}
	if topCentrePileCard.cardType == "Regular" {
		if playerCard.cardType == "Wild" {
			return true
		} else if playerCard.cardType == "Regular" {
			if playerCard.color == topCentrePileCard.color || playerCard.value == topCentrePileCard.value {
				return true
			}
		} else if playerCard.cardType == "Skip" || playerCard.cardType == "Reverse" || playerCard.cardType == "DrawTwo" {
			if playerCard.color == topCentrePileCard.color {
				return true
			}
		}
	}
	return false
}

func allowedPlayDeck(centrePile deck, playerHand deck) deck {
	/*
		Arguments:
		centrePile - cards where the players play their cards
		playerHand - deck consisting of the player's hand

		Returns:
		deck of cards that the player is "allowed/valid" to play
		returns and empty deck is no card is "allowed/playable"
	*/
	topCardFromPile := getTop(centrePile)
	allowedDeck := deck{} // contains the cards that can be played
	for _, c := range playerHand {
		if allowedPlayCard(c, topCardFromPile) {
			allowedDeck = append(allowedDeck, c)
		}
	}
	return allowedDeck
}

func nextPlayer(currentPlayerIndex int, allPlayers players) (int, string) {
	/*
		Arguments:
		currentPlayerIndex - integer keeping track of the index of the current player
		allPlayers - players obj (slice containing all player objects)

		Returns:
		int - index of the next player
		string - name of next player
	*/
	return (currentPlayerIndex + 1) % len(allPlayers), allPlayers[(currentPlayerIndex+1)%len(allPlayers)].name
}

func playerPlayCard(playerHand deck, c card, centrePile deck) (deck, deck) {
	/*
		Removes card c from the deck playerhand and places it on top of centrePile deck
		Arguments:
		playerHand - deck consisting of player's card
		c - card in the deck playerHand
		centrePile - deck consisting of all the cards where the players are playing their cards

		Returns:
		Updated playerHand deck
		updated centrePile deck
	*/

	for requiredCardIndex, requiredCard := range playerHand {
		if c == requiredCard {
			centrePile = append(centrePile, c)
			playerHand = removeCardFromDeck(requiredCardIndex, playerHand)
		}
	}
	return playerHand, centrePile
}

func playerTurnInput(p player, centrePile deck) (player, deck) {
	/*
		Takes a player and the pile of cards
		Asks the player which card to play (only from the ones that are allowed to)
		That card is moved from the player's hand to the centrepile
	*/
	d := p.hand
	playableHand := allowedPlayDeck(centrePile, d)
	if len(playableHand) == 0 {
	}
	printDeck(playableHand)
	var playCardIndex int
	fmt.Scanln(&playCardIndex)
	selectedCard := playableHand[playCardIndex]
	centrePile = addCardToDeck(selectedCard, centrePile)
	//Remove card from the player's hand

	for i, c := range d {
		if c == selectedCard {
			d = removeCardFromDeck(i, d)
			p.hand = d
		}
	}
	return p, centrePile
}
