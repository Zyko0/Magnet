package core

import (
	"math"
	"math/rand"

	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PlayerFallingVelocity      = 7.5
	PlayerSlidingVelocityAngle = 0.02
)

type Game struct {
	ticks uint64
	seed  float32

	Ring   *Ring
	Player *Player
}

func NewGame() *Game {
	return &Game{
		ticks: 0,
		seed:  rand.Float32(),

		Ring:   newRing(),
		Player: newPlayer(),
	}
}

func (g *Game) movePlayer() {
	x, y := ebiten.CursorPosition()
	cv := geom.Vec2{
		X: float32(x) - g.Ring.Center.X,
		Y: float32(y) - g.Ring.Center.Y,
	}
	cv.Normalize()
	ca := cv.Atan2()

	pv := geom.Vec2{
		X: g.Player.Position.X - g.Ring.Center.X,
		Y: g.Player.Position.Y - g.Ring.Center.Y,
	}
	pv.Normalize()
	pa := pv.Atan2()

	v := geom.Vec2{
		X: float32(x) - g.Player.Position.X,
		Y: float32(y) - g.Player.Position.Y,
	}
	v.Normalize()
	v.MulN(PlayerFallingVelocity)

	position := g.Player.Position
	position.Add(v)
	position.Add(g.Ring.getPlayerRingVelocity(g.Player))

	// Handle sliding
	var angle = pa
	if dist := position.DistanceTo(g.Ring.Center); dist > MaxPlayerDistance {
		g.Player.BonesSet = assets.BoneSetSliding

		// Note: hacky because bad at maths :D
		testa := (math.Pi + ca) - (math.Pi + pa)
		if math.Abs(float64(ca-pa)) <= PlayerSlidingVelocityAngle {
			angle = ca + math.Pi
		} else if testa < -math.Pi {
			angle = angle + math.Pi + PlayerSlidingVelocityAngle
		} else if testa > math.Pi {
			angle = angle + math.Pi - PlayerSlidingVelocityAngle
		} else if testa > 0 {
			angle = angle + math.Pi + PlayerSlidingVelocityAngle
		} else if testa < 0 {
			angle = angle + math.Pi - PlayerSlidingVelocityAngle
		}
		angle = float32(math.Mod(float64(angle), 2*math.Pi)) - math.Pi

		s, c := math.Sincos(float64(angle))
		position.X = float32(c) * MaxPlayerDistance
		position.Y = float32(s) * MaxPlayerDistance
		position.Add(g.Ring.Center)
		// Set rotation after boneset change
		g.Player.setRotation(cv.X, cv.Y, angle)
		// Set updated position
		g.Player.Position = position
		return
	}

	g.Player.BonesSet = assets.BoneSetFalling

	r := g.Player.Rotation
	r.Add(geom.Vec3{
		X: cv.X * 0.1,
		Y: cv.Y * 0.1,
		Z: angle * 0.1,
	})
	// Set rotation after boneset change
	g.Player.setRotation(r.X, r.Y, r.Z)
	// Set updated position
	g.Player.Position = position
}

func (g *Game) Update() {
	g.Ring.Update()

	// TODO: if cursor didn't change between last turn, don't move player
	g.movePlayer()

	g.ticks++
}
