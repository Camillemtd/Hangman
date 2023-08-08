package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid word=''")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validSate(t, "goodGuess", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("p")
	validSate(t, "badGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")
	validSate(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(2, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("z")
	validSate(t, "lost", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(2, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("a")
	validSate(t, "alreadyGuessed", g.State)
}

func validSate(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v'. got=%v", expectedState, actualState)
		return (false)
	}
	return true
}
