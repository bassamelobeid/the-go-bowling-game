package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rollMultipleTimes(game *Game, rolls int, pins int) {
	for i := 0; i < rolls; i++ {
		game.Roll(pins)
	}
}

func rollSpare(game *Game) {
	game.Roll(5)
	game.Roll(5)
}

func rollStrike(game *Game) {
	game.Roll(10)
}

func TestGutterGame(t *testing.T) {
	game := NewGame()
	rollMultipleTimes(game, 20, 0)
	assert.Equal(t, 0, game.GetScore())
}

func TestAllOnesGame(t *testing.T) {
	game := NewGame()
	rollMultipleTimes(game, 20, 1)
	assert.Equal(t, 20, game.GetScore())
}

func TestOneSpare(t *testing.T) {
	game := NewGame()
	rollSpare(game)
	game.Roll(3)
	rollMultipleTimes(game, 17, 0)
	assert.Equal(t, 16, game.GetScore())
}

func TestOneStrike(t *testing.T) {
	game := NewGame()
	rollStrike(game)
	game.Roll(3)
	game.Roll(4)
	rollMultipleTimes(game, 16, 0)
	assert.Equal(t, 24, game.GetScore())
}

func TestPerfectGame(t *testing.T) {
	game := NewGame()
	rollMultipleTimes(game, 12, 10)
	assert.Equal(t, 300, game.GetScore())
}
