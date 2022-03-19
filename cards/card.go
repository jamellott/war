package cards

import (
	"fmt"
	"strconv"
)

// Card represents a card from a standard 52-card deck.
type Card uint8

/// Suit represents the suit of a `Card`
type Suit uint8

/// Value represents a `Card`'s numerical value (such as `3`, `Jack`, `Ace`).
/// Face cards (including `Ace`) have named constants. Numeric cards can
/// be obtained by casting their numeric card to a Value.
/// `Ace` is considered to be high card.
type Value uint8

const (
	Clubs Suit = iota
	Diamonds
	Spades
	Hearts

	firstSuit = int(Clubs)
	suitCount = 4
)

const (
	Jack Value = 11 + iota
	Queen
	King
	Ace

	firstValue = 2
	valueCount = 13

	// helper constant, rounded up to power of 2
	suitModulo = 16
)

func NewCard(val Value, suit Suit) Card {
	return Card(uint8(val) + uint8(suit)*suitModulo)
}

func (c Card) Value() Value {
	return Value(c % suitModulo)
}

func (c Card) Suit() Suit {
	return Suit(c / suitModulo)
}

func (v Value) Name() string {
	switch v {
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return strconv.Itoa(int(v))
	}
}

func (s Suit) Name() string {
	switch s {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	default:
		return "unknown suit"
	}
}

/// Name returns the card's name in "`Value` of `Suit`" form.
func (c Card) Name() string {
	return fmt.Sprintf("%v of %v", c.Value().Name(), c.Suit().Name())
}

/// Compare returns -1, 0, or 1 depending on if the left card is
/// less, equivalent, or greater in `Value` to the right card.
func Compare(lhs, rhs Card) int {
	switch {
	case lhs.Value() < rhs.Value():
		return -1
	case lhs.Value() > rhs.Value():
		return 1
	default:
		return 0
	}
}
