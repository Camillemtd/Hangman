package hangman

import (
	"fmt"
)

func DrawWelcome() {
	fmt.Println(`
	
	__    __       ___      .__   __.   _______ .___  ___.      ___      .__   __. 
	|  |  |  |     /   \     |  \ |  |  /  _____||   \/   |     /   \     |  \ |  | 
	|  |__|  |    /  ^  \    |   \|  | |  |  __  |  \  /  |    /  ^  \    |   \|  | 
	|   __   |   /  /_\  \   |  .    | |  | |_ | |  |\/|  |   /  /_\  \   |  .    | 
	|  |  |  |  /  _____  \  |  |\   | |  |__| | |  |  |  |  /  _____  \  |  |\   | 
	|__|  |__| /__/     \__\ |__| \__|  \______| |__|  |__| /__/     \__\ |__| \__| 
																					
	
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		 ____
		|/   |
		|   (_)
		|   /|\
		|    |
		|   | |
		|
		|_____
		`
	case 1:
		draw = `
		 ____
		|/   |
		|   (_)
		|   \|/
		|    |
		|   / \
		|
		|_____
		`
	case 2:
		draw = `
		 ____
		|/   |
		|   (_)
		|   \|/
		|    |
		|   / 
		|
		|_____
		`
	case 3:
		draw = `
		 ____
		|/   |
		|   (_)
		|   \|/
		|    |
		|    
		|
		|_____
		`
	case 4:
		draw = `
		 ____
		|/   |
		|   (_)
		|   \|
		|    |
		|    
		|
		|_____
		`
	case 5:
		draw = `
		 ____
		|/   |
		|   (_)
		|    |
		|    |    
		|    
		|
		|_____
		`
	case 6:
		draw = `
		 ____
		|/   |
		|   (_)
		|    
		|    
		|    
		|
		|_____
		`
	case 7:
		draw = `
		 ____
		|/   |
		|   
		|    
		|    
		|    
		|
		|_____
		`
	case 8:
		draw = `
		
		`
	}
	fmt.Println(draw)

}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word \n", guess)
	case "lost":
		fmt.Print("You lost:! The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON! The word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
