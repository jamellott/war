package cards

import "math/rand"

/// Deck represents a deck of cards that may be held by a player.
type Deck struct {
	cards []Card
}

/// NewDeck returns a standard 52 card deck.
func NewDeck() *Deck {
	deck := make([]Card, 0, suitCount*valueCount)
	for s := 0; s < suitCount; s++ {
		for v := 0; v < valueCount; v++ {
			deck = append(deck, NewCard(Value(firstValue+v), Suit(firstSuit+s)))
		}
	}

	return &Deck{
		cards: deck,
	}
}

/// Split splits a deck evenly into two half sized decks. Panics if Len()
/// is not even.
func (d *Deck) Split() [2]Deck {
	if d.Len()%2 != 0 {
		panic("tried to split unevenly sized deck")
	}

	// Clone underlying slices
	midpoint := d.Len() / 2
	d1 := Deck{cards: append([]Card{}, d.cards[:midpoint]...)}
	d2 := Deck{cards: append([]Card{}, d.cards[midpoint:]...)}

	return [2]Deck{d1, d2}
}

/// Deal takes a card off the top of the deck and returns it.
func (d *Deck) Deal() Card {
	if d.Len() == 0 {
		panic("dealt from empty deck")
	}

	card := d.cards[0]
	d.cards = d.cards[1:]

	return card
}

/// Len returns the number of cards remaining in the deck.
func (d *Deck) Len() int {
	return len(d.cards)
}

/// Shuffle is a helper for shuffling based on the math/rand library.
func (d *Deck) Shuffle() {
	rand.Shuffle(d.Len(), func(i, j int) {
		tmp := d.cards[i]
		d.cards[i] = d.cards[j]
		d.cards[j] = tmp
	})
}

/// Append places the given cards at the bottom of the deck.
func (d *Deck) Append(cards ...Card) {
	d.cards = append(d.cards, cards...)
}
