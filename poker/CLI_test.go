package poker_test

import (
	"learn_go_with_tests/poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("Record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PLayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("Record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PLayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
