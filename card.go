package main

import "fmt"

// import "fmt"

type card struct {
	cardType string // whether the card is regular or special
	value    string
	color    string
}

//type1 cards - wild and wildDrawFour
//type2 cards - reverse, skip and drawTwo
//type3 cards - regular

func createType1Card(cardType string) card {
	var card card
	card.cardType = cardType
	return card
}

func createType2Card(cardType string, color string) card {
	var card card
	card.cardType = cardType
	card.color = color
	return card
}

func createType3Card(value string, color string) card {
	var card card
	card.value = value
	card.color = color
	return card
}

func printCard(card card) {
	if card.cardType == "Wild" || card.cardType == "WildDrawFour" {
		fmt.Println(card.cardType)
	} else if card.cardType == "Reverse" || card.cardType == "Skip" || card.cardType == "DrawTwo" {
		fmt.Printf("%v %v\n", card.color, card.cardType)
	} else {
		fmt.Printf("%v %v\n", card.value, card.color)
	}
}
