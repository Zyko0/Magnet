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
	ticks  uint64
	seed   float32
	cursor geom.Vec2

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

func (g *Game) movePlayer(input bool) {
	var (
		cv, pv, dir geom.Vec2
		ca, pa      float32
	)

	// Normalize cursor vector
	x, y := g.cursor.X, g.cursor.Y
	cv.X = float32(x) - g.Ring.Center.X
	cv.Y = float32(y) - g.Ring.Center.Y
	cv.Normalize()
	ca = cv.Atan2()
	// Normalize player vector
	pv.X = g.Player.Position.X - g.Ring.Center.X
	pv.Y = g.Player.Position.Y - g.Ring.Center.Y
	pv.Normalize()
	pa = pv.Atan2()
	// Player to cursor velocity
	dir.X = float32(x) - g.Player.Position.X
	dir.Y = float32(y) - g.Player.Position.Y
	// Below is required to avoid flickering and continuous position updates
	const updateThreshold = 4
	if dir.Length() < updateThreshold {
		dir.X, dir.Y = 0, 0
	}

	dir.Normalize()

	position := dir
	position.MulN(PlayerFallingVelocity)
	position.Add(g.Player.Position)
	position.Add(g.Ring.getPlayerRingVelocity(g.Player))

	// Handle sliding
	var angle = pa
	if dist := position.DistanceTo(g.Ring.Center); dist > MaxPlayerDistance {
		g.Player.BonesSet = assets.BoneSetSliding

		// Note: hacky because bad at maths
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

	r := g.Player.Rotation
	switch g.Player.BonesSet {
	case assets.BoneSetFalling:
		r.Add(geom.Vec3{
			X: cv.X * 0.125,
			Y: cv.Y * 0.125,
			Z: angle * 0.1,
		})
	case assets.BoneSetDashing:
		// Update player angle on dash only if direction changed
		r.X = dir.X
		r.Y = dir.Y
		r.Z = dir.Atan2()
	}
	// Set rotation after boneset change
	g.Player.setRotation(r.X, r.Y, r.Z)
	g.Player.Position = position
}

func (g *Game) Update() {
	g.Ring.Update()
	g.Player.Update()

	x, y := ebiten.CursorPosition()
	cursor := geom.Vec2{
		X: float32(x),
		Y: float32(y),
	}
	input := !g.cursor.Equals(cursor)
	g.cursor = cursor
	// TODO: if cursor didn't change between last turn, don't move player
	g.movePlayer(input)

	g.ticks++
}
