package main

import (
	"fmt"
	"github.com/jamellott/war/game"
	"math/rand"
)

func main() {
    var seed int64
    fmt.Printf("Enter game seed: ")
    fmt.Scanf("%v\n", &seed)
	rand.Seed(seed)

	DisplayGame(game.RunGame())
}

func DisplayGame(report []game.WarReport) {
	for warIndex := range report {
		fmt.Printf("\n\nWar %v\n==========\n", warIndex+1)
		war := &report[warIndex]

		for battleIndex := range war.Battles {
			fmt.Printf("\nBattle %v\n----------\n", battleIndex+1)
			battle := &war.Battles[battleIndex]
			DisplayBattle(battle)
		}

		if war.EndsGame {
			switch war.Winner {
			case game.TieWin:
				fmt.Println("Neither player has enough cards to continue playing")
				fmt.Println("Game Tie!")
			case game.Player1Win:
				fmt.Println("Player 2 does not have enough cards to continue playing")
				fmt.Println("Game Win! Player 1")
			case game.Player2Win:
				fmt.Println("Player 1 does not have enough cards to continue playing")
				fmt.Println("Game Win! Player 2")
			}

			return
		}

		switch war.Winner {
		case game.Player1Win:
			fmt.Println("Player 1 wins the war!")
		case game.Player2Win:
			fmt.Println("Player 2 wins the war!")
		}

		fmt.Println("Total winnings:")
		for _, card := range war.Pot {
			fmt.Printf("- %v\n", card.Name())
		}

		fmt.Printf("Player %v now has %v card(s)\n", 1, war.RemainingCards[0])
		fmt.Printf("Player %v now has %v card(s)\n", 2, war.RemainingCards[1])
	}
}

func DisplayBattle(report *game.BattleReport) {
	fmt.Printf("Player %v plays %v\n", 1, report.CardsPlayed[0].Name())
	fmt.Printf("Player %v plays %v\n", 2, report.CardsPlayed[1].Name())

	switch report.Winner {
	case game.TieWin:
		fmt.Println("Battle inconclusive. The war continues...")
	case game.Player1Win:
		fmt.Println("Player 1 wins the battle!")
	case game.Player2Win:
		fmt.Println("Player 2 wins the battle!")
	}
}
