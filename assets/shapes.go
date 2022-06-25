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
	ShapeIndexPortal = iota
	ShapeIndexLine
	ShapeIndexCross
	ShapeIndexEmptyLine
	ShapeIndexEmptyCross
	ShapeIndexHalfQuad
	ShapeIndexThinEmptyCross
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
	trianglesLine = []*geom.Triangle{
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
	}
	trianglesCross = []*geom.Triangle{
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
	trianglesEmptyLine = []*geom.Triangle{
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
				X: 1,
				Y: 0.3,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0,
				Y: 0.3,
			},
			C: geom.Vec2{
				X: 1,
				Y: 0.3,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.7,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0.7,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.7,
			},
			B: geom.Vec2{
				X: 0,
				Y: 1,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
	}
	trianglesEmptyCross = []*geom.Triangle{
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0.35,
				Y: 0,
			},
			C: geom.Vec2{
				X: 0.35,
				Y: 0.35,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0,
				Y: 0.35,
			},
			C: geom.Vec2{
				X: 0.35,
				Y: 0.35,
			},
		},
		{
			A: geom.Vec2{
				X: 0.65,
				Y: 0,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0,
			},
			C: geom.Vec2{
				X: 1,
				Y: 0.35,
			},
		},
		{
			A: geom.Vec2{
				X: 0.65,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0.65,
				Y: 0.35,
			},
			C: geom.Vec2{
				X: 1,
				Y: 0.35,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.65,
			},
			B: geom.Vec2{
				X: 0.35,
				Y: 0.65,
			},
			C: geom.Vec2{
				X: 0.35,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.65,
			},
			B: geom.Vec2{
				X: 0,
				Y: 1,
			},
			C: geom.Vec2{
				X: 0.35,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0.65,
				Y: 0.65,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0.65,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0.65,
				Y: 0.65,
			},
			B: geom.Vec2{
				X: 0.65,
				Y: 1,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
	}
	trianglesHalfQuad = []*geom.Triangle{
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0.5,
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
				X: 0.5,
				Y: 0,
			},
			C: geom.Vec2{
				X: 0.5,
				Y: 1,
			},
		},
	}
	trianglesThinEmptyCross = []*geom.Triangle{
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0.375,
				Y: 0,
			},
			C: geom.Vec2{
				X: 0.375,
				Y: 0.375,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0,
				Y: 0.375,
			},
			C: geom.Vec2{
				X: 0.375,
				Y: 0.375,
			},
		},
		{
			A: geom.Vec2{
				X: 0.625,
				Y: 0,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0,
			},
			C: geom.Vec2{
				X: 1,
				Y: 0.375,
			},
		},
		{
			A: geom.Vec2{
				X: 0.625,
				Y: 0,
			},
			B: geom.Vec2{
				X: 0.625,
				Y: 0.375,
			},
			C: geom.Vec2{
				X: 1,
				Y: 0.375,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.625,
			},
			B: geom.Vec2{
				X: 0.375,
				Y: 0.625,
			},
			C: geom.Vec2{
				X: 0.375,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0,
				Y: 0.625,
			},
			B: geom.Vec2{
				X: 0,
				Y: 1,
			},
			C: geom.Vec2{
				X: 0.375,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0.625,
				Y: 0.625,
			},
			B: geom.Vec2{
				X: 1,
				Y: 0.625,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
		{
			A: geom.Vec2{
				X: 0.625,
				Y: 0.625,
			},
			B: geom.Vec2{
				X: 0.625,
				Y: 1,
			},
			C: geom.Vec2{
				X: 1,
				Y: 1,
			},
		},
	}

	TriangleShapes = [][]*geom.Triangle{
		ShapeIndexPortal:         trianglesPortal,
		ShapeIndexLine:           trianglesLine,
		ShapeIndexCross:          trianglesCross,
		ShapeIndexEmptyLine:      trianglesEmptyLine,
		ShapeIndexEmptyCross:     trianglesEmptyCross,
		ShapeIndexHalfQuad:       trianglesHalfQuad,
		ShapeIndexThinEmptyCross: trianglesThinEmptyCross,
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
