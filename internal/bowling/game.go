package bowling

type Game struct {
	rolls       []int
	currentRoll int
}

func NewGame() *Game {
	return &Game{
		rolls:       make([]int, 21),
		currentRoll: 0,
	}
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) GetScore() int {
	score := 0
	firstInFrame := 0
	for frame := 0; frame < 10; frame++ {
		if g.isSpare(firstInFrame) {
			score += 10 + g.rolls[firstInFrame+2]
			firstInFrame += 2
		} else if g.isStrike(firstInFrame) {
			score += 10 + g.rolls[firstInFrame+1] + g.rolls[firstInFrame+2]
			firstInFrame++
		} else {
			score += g.rolls[firstInFrame] + g.rolls[firstInFrame+1]
			firstInFrame += 2
		}
	}
	return score
}

func (g *Game) isSpare(firstInFrame int) bool {
	return g.rolls[firstInFrame]+g.rolls[firstInFrame+1] == 10
}

func (g *Game) isStrike(firstInFrame int) bool {
	return g.rolls[firstInFrame] == 10
}
