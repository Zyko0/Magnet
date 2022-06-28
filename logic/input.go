package logic

import (
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	Touches []ebiten.TouchID
	Cursor  geom.Vec2

	IntentDash bool
)

func UpdateInputs() {
	Touches = ebiten.AppendTouchIDs(Touches[:0])

	if len(Touches) > 0 {
		touch := Touches[len(Touches)-1]
		IntentDash = inpututil.TouchPressDuration(touch) > 5 // ~80ms

		x, y := ebiten.TouchPosition(touch)
		Cursor.X, Cursor.Y = float32(x), float32(y)
	} else {
		IntentDash = ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

		x, y := ebiten.CursorPosition()
		// Note: do not update cursor if we're on mobile, it would reset it to 0,0
		if x+y > 0 {
			Cursor.X, Cursor.Y = float32(x), float32(y)
		}
	}
}
