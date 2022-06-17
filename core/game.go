package core

type Game struct {
	ticks uint64
}

func NewGame() *Game {
	return &Game{
		ticks: 0,
	}
}

func (g *Game) Update() error {
	g.ticks++

	return nil
}
