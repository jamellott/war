package game

import "github.com/jamellott/war/cards"

/// RoundWinner is an enum representing who won a specific event
type RoundWinner int

const (
	Player1Win RoundWinner = iota
	Player2Win
	TieWin
)

/// BattleReport stores a record of a single battle within a war.
type BattleReport struct {
	/// CardsPlayed lists the card played by each player, keyed by index
	CardsPlayed []cards.Card
	Winner      RoundWinner
}

/// WarReport stores a record of a single war within a game
type WarReport struct {
	Battles []BattleReport

	/// Pot lists the cards that were won by the eventual winner
	Pot []cards.Card

	/// EndsGame lists if this was the final war of the game. If it is true,
	/// Winner also represents the winner of the game overall.
	EndsGame bool
	Winner   RoundWinner

	/// RemainingCards lists the number of cards each player has in their hand
	/// following completion of this war.
	RemainingCards []int
}

// checkEnd returns whether the game is over, and if so, who won the game.
func checkEnd(hands []cards.Deck) (result bool, winner RoundWinner) {
	switch {
	case hands[0].Len() == 0 && hands[1].Len() == 0:
		return true, TieWin
	case hands[1].Len() == 0:
		return true, Player1Win
	case hands[0].Len() == 0:
		return true, Player2Win
	}

	return false, TieWin
}

// runBattle simulates a single battle within a war. Each player will play
// a card, then the winner of the battle is the player with the higher valued card.
func runBattle(hands []cards.Deck) BattleReport {
	cardsPlayed := []cards.Card{hands[0].Deal(), hands[1].Deal()}

	result := cards.Compare(cardsPlayed[0], cardsPlayed[1])

	switch result {
	case 0:
		// battle inconclusive
		return BattleReport{
			Winner:      TieWin,
			CardsPlayed: cardsPlayed,
		}
	case 1:
		return BattleReport{
			Winner:      Player1Win,
			CardsPlayed: cardsPlayed,
		}
	case -1:
		return BattleReport{
			Winner:      Player2Win,
			CardsPlayed: cardsPlayed,
		}
	}

	panic("cards.Compare returned unknown value")
}

// The number of cards played each time a war continues
const cardsPerWarBattle = 3

// runWar simulates a single war within a game. The two players will run
// battles against each other repeatedly until one of them wins a battle
// or runs out of cards. In between each battle, 3 cards are drawn from
// each players deck and placed inside a winnings pot, along with every
// card that was played. The first player to win a battle receives all cards inside the pot.
func runWar(hands []cards.Deck) WarReport {
	battles := []BattleReport{runBattle(hands)}
	pot := append([]cards.Card{}, battles[0].CardsPlayed...)

	if gameOver, winner := checkEnd(hands); gameOver {
		return WarReport{
			Winner:         winner,
			Battles:        battles,
			Pot:            pot,
			EndsGame:       true,
			RemainingCards: []int{hands[0].Len(), hands[1].Len()},
		}
	}

	for battles[len(battles)-1].Winner == TieWin {
		for i := 0; i < cardsPerWarBattle; i++ {
			pot = append(pot, hands[0].Deal(), hands[1].Deal())
			if gameOver, winner := checkEnd(hands); gameOver {
				return WarReport{
					Winner:         winner,
					Battles:        battles,
					Pot:            pot,
					EndsGame:       true,
					RemainingCards: []int{hands[0].Len(), hands[1].Len()},
				}
			}
		}

		battle := runBattle(hands)
		battles = append(battles, battle)
		pot = append(pot, battle.CardsPlayed...)
	}

	// award winnings
	switch battles[len(battles)-1].Winner {
	case Player1Win:
		hands[0].Append(pot...)
	case Player2Win:
		hands[1].Append(pot...)
	}

	return WarReport{
		Winner:         battles[len(battles)-1].Winner,
		Battles:        battles,
		Pot:            pot,
		EndsGame:       hands[0].Len() == 0 || hands[1].Len() == 0,
		RemainingCards: []int{hands[0].Len(), hands[1].Len()},
	}
}

/// StartGame simulates an entire game of wars and returns a series of reports
/// detailing what happened in each war.
func StartGame() <-chan WarReport {
	fullDeck := cards.NewDeck()
	fullDeck.Shuffle()
	hands := fullDeck.Split()

    reports := make(chan WarReport)
    go func() {
        for {
            war := runWar(hands[:])
            reports <- war

            if war.EndsGame {
                close(reports)
                return
            }
        }
    }()

	return reports
}
