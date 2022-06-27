package logic

import (
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Touches []ebiten.TouchID
	Cursor  geom.Vec2

	IntentDash bool
)

func UpdateInputs() {
	Touches = ebiten.AppendTouchIDs(Touches[:0])

	if len(Touches) > 0 {
		IntentDash = len(Touches) >= 2

		x, y := ebiten.TouchPosition(Touches[0])
		Cursor.X, Cursor.Y = float32(x), float32(y)
	} else {
		IntentDash = ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

		x, y := ebiten.CursorPosition()
		Cursor.X, Cursor.Y = float32(x), float32(y)
	}
}
