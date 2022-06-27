package ui

import (
	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	splashDisplayDuration = logic.TPS * 4
)

type SplashView struct {
	ticks  uint
	active bool

	colorm ebiten.ColorM
}

func NewSplashView() *SplashView {
	return &SplashView{
		ticks:  0,
		active: true,

		colorm: ebiten.ColorM{},
	}
}

func (sv *SplashView) Active() bool {
	return sv.active
}

func (sv *SplashView) Update() {
	sv.ticks++

	if sv.ticks > splashDisplayDuration {
		sv.active = false
		return
	}
	if len(inpututil.AppendPressedKeys(nil)) > 0 || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || len(logic.Touches) > 0 {
		sv.active = false
	}

	d := float64(sv.ticks) / float64(splashDisplayDuration)
	sc := (-(d * d) + d) * 4
	sv.colorm.Reset()
	sv.colorm.Scale(sc, sc, sc, 1.)
}

func (sv *SplashView) Draw(screen *ebiten.Image) {
	screen.DrawImage(assets.SplashImage, &ebiten.DrawImageOptions{
		Filter: ebiten.FilterLinear,
		ColorM: sv.colorm,
	})
}
