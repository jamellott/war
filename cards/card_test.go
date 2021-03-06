package cards

import "testing"

// TestCard calls cards.NewCard, cards.Card.Value, and cards.Card.Suit to
// check that card suits and values are packed and unpacked correctly.
func TestCard(t *testing.T) {
	card := NewCard(Value(7), Spades)

	if card.Value() != Value(7) {
		t.Fatalf("Expected \"7 of Spades\", got \"%v\"", card.Name())
	}
}

// TestCardName calls cards.Card.Name to check that a card name is constructed
// from its value and suit correctly.
func TestCardName(t *testing.T) {
	{
		card := NewCard(Ace, Diamonds)

		if card.Name() != "Ace of Diamonds" {
			t.Fatalf("Expected \"Ace of Diamonds\", got \"%v\"", card.Name())
		}
	}

	{
		card := NewCard(3, Clubs)

		if card.Name() != "3 of Clubs" {
			t.Fatalf("Expected \"3 of Clubs\", got \"%v\"", card.Name())
		}
	}
}

// TestCompare calls cards.Compare to ensure that cards, including
// Aces, compare correctly.
func TestCompare(t *testing.T) {
	{
		c1 := NewCard(Value(2), Clubs)
		c2 := NewCard(Value(4), Clubs)

		if result := Compare(c1, c2); result != -1 {
			t.Errorf("Expected comparison of \"2 of Clubs\" and \"4 of Clubs\" to return -1, got %v", result)
		}
	}

	{
		c1 := NewCard(Value(4), Clubs)
		c2 := NewCard(Value(2), Clubs)

		if result := Compare(c1, c2); result != 1 {
			t.Errorf("Expected comparison of \"4 of Clubs\" and \"2 of Clubs\" to return 1, got %v", result)
		}
	}

	{
		c1 := NewCard(Ace, Clubs)
		c2 := NewCard(Value(4), Diamonds)

		if result := Compare(c1, c2); result != 1 {
			t.Errorf("Expected comparison of \"Ace of Clubs\" and \"4 of Diamonds\" to return 1, got %v", result)
		}
	}

	{
		c1 := NewCard(Queen, Hearts)
		c2 := NewCard(Queen, Diamonds)

		if result := Compare(c1, c2); result != 0 {
			t.Errorf("Expected comparison of \"Queen of Hearts\" and \"Queen of Diamonds\" to return 0, got %v", result)
		}
	}

	{
		c1 := NewCard(Jack, Spades)
		c2 := NewCard(Value(8), Diamonds)

		if result := Compare(c1, c2); result != 1 {
			t.Errorf("Expected comparison of \"Jack of Spades\" and \"8 of Diamonds\" to return 1, got %v", result)
		}
	}
}

// TestSplit calls cards.Deck.Split() and ensures that the two returned
// decks are of equal size.
func TestSplit(t *testing.T) {
	d := NewDeck()

	size := d.Len()
	decks := d.Split()
	if decks[0].Len() != size/2 && decks[1].Len() != size/2 {
		t.Fatalf("Expected decks to be split in half. Expected: 52,26,26 actual %v,%v,%v", size, decks[0].Len(), decks[1].Len())
	}
}

// TestDeal calls cards.Deck.Deal() and ensures that the number of cards
// in the deck are decremented.
func TestDeal(t *testing.T) {
	d := NewDeck()

	size := d.Len()
	_ = d.Deal()

	if d.Len() != size-1 {
		t.Fatalf("Expected deck size %v, actual %v", size-1, size)
	}
}
