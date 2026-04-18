package blackjack


// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
        case "jack", "queen", "king", "ten":
        	return 10
		case "one":
        	return 1
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
    	default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Value := ParseCard(card1)
    card2Value := ParseCard(card2)
    totalCardsValue := card1Value + card2Value
    dealerCardValue := ParseCard(dealerCard)
	switch {
        case totalCardsValue > 21:
        	return "P"
        case totalCardsValue == 21 && dealerCardValue < 10:
        	return "W"
    	case totalCardsValue == 21 && dealerCardValue >= 10:
        	return "S"
        case totalCardsValue >= 17 && totalCardsValue <= 20:
        	return "S"
        case totalCardsValue >= 12 && totalCardsValue <= 16 && dealerCardValue < 7:
        	return "S"
        case totalCardsValue >= 12 && totalCardsValue <= 16 && dealerCardValue >= 7:
        	return "H"
        // case totalCardsValue <= 11:
        default:
        	return "H"
        
        
    }
}
