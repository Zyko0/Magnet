package assets

import (
	"github.com/Zyko0/Magnet/graphics"
	"github.com/Zyko0/Magnet/logic"
	"github.com/Zyko0/Magnet/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ShapeSize    = logic.ScreenHeight
	ShapeOffsetX = logic.ScreenWidth/2 - ShapeSize/2
)

var (
	ShapeCircleMaskImage = ebiten.NewImage(ShapeSize, ShapeSize)
)

const (
	ShapeIndexTrianglePortal = iota
	ShapeIndexTrianglePlus
)

var (
	trianglesPortal = []*geom.Triangle{
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0,
			},
			C: geom.Vec2{
				X: 0,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 1,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
	}
	trianglesPlus = []*geom.Triangle{
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.4,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0.4,
			},
			C: geom.Vec2{
				X: 0,
				Y: 0.6,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.6,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0.4,
			},
			C: geom.Vec2{
				X: 1,
				Y: 0.6,
			},
		},
		{
			A: geom.Vec2{
				X: 0.4,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0.4,
				Y: 1,
			},
			C: geom.Vec2{
				X: 0.6,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0.6,
				Y: 1,
			},
			B: geom.Vec2{
				X: 0.6,
				Y: 0,
			},
			C: geom.Vec2{
				X: 0.4,
				Y: 0,
			},
		},
	}

	TriangleShapes = [][]*geom.Triangle{
		ShapeIndexTrianglePortal: trianglesPortal,
		ShapeIndexTrianglePlus:   trianglesPlus,
	}
)

func init() {
	// TODO: Draw a circle in the global image
	vertices, indices := graphics.AppendQuadVerticesIndices(nil, nil,
		0, 0, ShapeSize, ShapeSize,
		1, 1, 1, 1, 0,
	)
	ShapeCircleMaskImage.DrawTrianglesShader(vertices, indices, CircleShader, nil)

	for i := range TriangleShapes {
		for _, t := range TriangleShapes[i] {
			t.A.MulN(ShapeSize)
			t.B.MulN(ShapeSize)
			t.C.MulN(ShapeSize)
			t.Translate(ShapeOffsetX, 0)
		}
	}
}