package blackjack

import "strings"

// ParseCard returns the value of the given card
func ParseCard(card string) int {
	card = strings.ToLower(card) // Convert card name to lowercase
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn determines the action based on the given cards
func FirstTurn(card1, card2, dealerCard string) string {
	card1 = strings.ToLower(card1) // Convert card names to lowercase
	card2 = strings.ToLower(card2)
	dealerCard = strings.ToLower(dealerCard)

	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	handValue := value1 + value2

	// Check for pairs of aces
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Check for Blackjack
	if handValue == 21 {
		if dealerValue == 11 || dealerValue == 10 {
			return "S"
		}
		return "W"
	}

	// Check for hand values within [17, 20]
	if handValue >= 17 {
		return "S"
	}

	// Check for hand values within [12, 16]
	if handValue >= 12 && handValue <= 16 {
		if dealerValue >= 7 {
			return "H"
		}
		return "S"
	}

	// Hand value is 11 or lower
	return "H"
}
