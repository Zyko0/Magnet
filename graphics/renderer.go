package graphics

import "github.com/hajimehoshi/ebiten/v2"

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(screen *ebiten.Image, options *GameDrawOptions) {
}
