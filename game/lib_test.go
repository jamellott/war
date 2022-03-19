package game

import (
	"github.com/jamellott/war/cards"
	"testing"
)

// testCardName compares a card to a given name
func testCardName(t *testing.T, expected string, card cards.Card) {
	if name := card.Name(); name != expected {
		t.Errorf("Expected %v, actual %v", expected, name)
	}
}

// Test_runBattle tests that a battle finds the correct winner
func Test_runBattle(t *testing.T) {
	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Ace, cards.Diamonds),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(2), cards.Diamonds),
			}),
		}

		report := runBattle(hands)

		if report.Winner != Player1Win {
			t.Errorf("Expected %v, actual %v", Player1Win, report.Winner)
		}

		testCardName(t, "Ace of Diamonds", report.CardsPlayed[0])
		testCardName(t, "2 of Diamonds", report.CardsPlayed[1])
	}

	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Queen, cards.Spades),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Queen, cards.Clubs),
			}),
		}

		report := runBattle(hands)

		if report.Winner != TieWin {
			t.Errorf("Expected %v, actual %v", TieWin, report.Winner)
		}

		testCardName(t, "Queen of Spades", report.CardsPlayed[0])
		testCardName(t, "Queen of Clubs", report.CardsPlayed[1])
	}

	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Clubs),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(10), cards.Hearts),
			}),
		}

		report := runBattle(hands)

		if report.Winner != Player2Win {
			t.Errorf("Expected %v, actual %v", Player2Win, report.Winner)
		}

		testCardName(t, "8 of Clubs", report.CardsPlayed[0])
		testCardName(t, "10 of Hearts", report.CardsPlayed[1])
	}
}

// Test_runWar_EndsGame tests that runWar will end a game upon tie or
// a player winning
func Test_runWar_EndsGame(t *testing.T) {
	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Clubs),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(10), cards.Hearts),
				cards.NewCard(cards.Value(2), cards.Hearts),
			}),
		}

		report := runWar(hands)

		if !report.EndsGame {
			t.Error("Expected war to end game")
		}

		if report.Winner != Player2Win {
			t.Errorf("Expected %v, actual %v", Player2Win, report.Winner)
		}

		if len(report.Battles) != 1 {
			t.Errorf("Expected %v, actual %v", 1, len(report.Battles))
		}

		testCardName(t, "8 of Clubs", report.Pot[0])
		testCardName(t, "10 of Hearts", report.Pot[1])
	}

	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Clubs),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Spades),
			}),
		}
		report := runWar(hands)

		if !report.EndsGame {
			t.Error("Expected war to end game")
		}

		if report.Winner != TieWin {
			t.Errorf("Expected %v, actual %v", TieWin, report.Winner)
		}

		if len(report.Battles) != 1 {
			t.Errorf("Expected %v, actual %v", 1, len(report.Battles))
		}

		testCardName(t, "8 of Clubs", report.Pot[0])
		testCardName(t, "8 of Spades", report.Pot[1])
	}
}

// Test_runWar_ContinuesGame tests that runWar will not end the game
// after a battle if the game is not over yet
func Test_runWar_ContinuesGame(t *testing.T) {
	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Clubs),
				cards.NewCard(cards.Value(9), cards.Clubs),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(10), cards.Hearts),
				cards.NewCard(cards.Value(2), cards.Hearts),
			}),
		}

		report := runWar(hands)

		if report.EndsGame {
			t.Error("Expected war to continue game")
		}

		if report.Winner != Player2Win {
			t.Errorf("Expected %v, actual %v", Player2Win, report.Winner)
		}

		if len(report.Battles) != 1 {
			t.Errorf("Expected %v, actual %v", 1, len(report.Battles))
		}

		testCardName(t, "8 of Clubs", report.Pot[0])
		testCardName(t, "10 of Hearts", report.Pot[1])

		if report.RemainingCards[0] != 1 {
			t.Errorf("Expected %v, actual %v", 1, report.RemainingCards[0])
		}

		if report.RemainingCards[1] != 3 {
			t.Errorf("Expected %v, actual %v", 3, report.RemainingCards[1])
		}
	}
}

// Test_runWar_MultiBattle tests that runWar will evaluate multi-battle wars
// and assign cards to the victor
func Test_runWar_MultiBattle(t *testing.T) {
	{
		hands := []cards.Deck{
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Clubs),
				cards.NewCard(cards.Value(9), cards.Clubs),
				cards.NewCard(cards.Value(9), cards.Clubs),
				cards.NewCard(cards.Value(9), cards.Clubs),
				cards.NewCard(cards.Value(6), cards.Clubs),
				cards.NewCard(cards.Value(5), cards.Clubs),
			}),
			*cards.NewDeckFromCards([]cards.Card{
				cards.NewCard(cards.Value(8), cards.Hearts),
				cards.NewCard(cards.Value(3), cards.Hearts),
				cards.NewCard(cards.Value(3), cards.Hearts),
				cards.NewCard(cards.Value(3), cards.Hearts),
				cards.NewCard(cards.Value(2), cards.Hearts),
				cards.NewCard(cards.Value(1), cards.Hearts),
			}),
		}

		report := runWar(hands)

		if report.EndsGame {
			t.Error("Expected war to continue game")
		}

		if report.Winner != Player1Win {
			t.Errorf("Expected %v, actual %v", Player2Win, report.Winner)
		}

		if len(report.Battles) != 2 {
			t.Errorf("Expected %v, actual %v", 1, len(report.Battles))
		}

		testCardName(t, "8 of Clubs", report.Pot[0])
		testCardName(t, "8 of Hearts", report.Pot[1])
		testCardName(t, "9 of Clubs", report.Pot[2])
		testCardName(t, "3 of Hearts", report.Pot[3])
		testCardName(t, "9 of Clubs", report.Pot[4])
		testCardName(t, "3 of Hearts", report.Pot[5])
		testCardName(t, "9 of Clubs", report.Pot[6])
		testCardName(t, "3 of Hearts", report.Pot[7])
		testCardName(t, "6 of Clubs", report.Pot[8])
		testCardName(t, "2 of Hearts", report.Pot[9])

		if report.RemainingCards[0] != 11 {
			t.Errorf("Expected %v, actual %v", 11, report.RemainingCards[0])
		}

		if report.RemainingCards[1] != 1 {
			t.Errorf("Expected %v, actual %v", 1, report.RemainingCards[1])
		}
	}
}
