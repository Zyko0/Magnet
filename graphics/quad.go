package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	brushImage = ebiten.NewImage(1, 1)
)

func init() {
	brushImage.Fill(color.RGBA{255, 255, 255, 255})
}

var (
	boxIndices = [6]uint16{0, 1, 2, 1, 2, 3}
)

func AppendQuadVerticesIndices(vertices []ebiten.Vertex, indices []uint16, x, y, w, h, r, g, b, a float32, index int) ([]ebiten.Vertex, []uint16) {
	vertices = append(vertices, []ebiten.Vertex{
		{
			DstX:   x,
			DstY:   y,
			SrcX:   0,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x + w,
			DstY:   y,
			SrcX:   1,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x,
			DstY:   y + h,
			SrcX:   0,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x + w,
			DstY:   y + h,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}...)

	indiceCursor := uint16(index * 4)
	indices = append(indices, []uint16{
		boxIndices[0] + indiceCursor,
		boxIndices[1] + indiceCursor,
		boxIndices[2] + indiceCursor,
		boxIndices[3] + indiceCursor,
		boxIndices[4] + indiceCursor,
		boxIndices[5] + indiceCursor,
	}...)

	return vertices, indices
}

func DrawRect(dst *ebiten.Image, x, y, width, height float32, r, g, b, a float32) {
	vertices, indices := AppendQuadVerticesIndices(nil, nil, x, y, width, height, r, g, b, a, 0)
	dst.DrawTriangles(vertices, indices, brushImage, &ebiten.DrawTrianglesOptions{})
}

var (
	borderBoxIndices = []uint16{0, 2, 4, 2, 4, 6, 1, 3, 5, 3, 5, 7}
)

func DrawRectBorder(dst *ebiten.Image, x, y, width, height, borderWidth, r, g, b, a float32) {
	dst.DrawTriangles([]ebiten.Vertex{
		{
			DstX:   x,
			DstY:   y,
			SrcX:   0,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX: x + borderWidth,
			DstY: y + borderWidth,
			SrcX: 0,
			SrcY: 0,
		},
		{
			DstX:   x + width,
			DstY:   y,
			SrcX:   1,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX: x + width - borderWidth,
			DstY: y + borderWidth,
			SrcX: 1,
			SrcY: 0,
		},
		{
			DstX:   x,
			DstY:   y + height,
			SrcX:   0,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX: x + borderWidth,
			DstY: y + height - borderWidth,
			SrcX: 0,
			SrcY: 1,
		},
		{
			DstX:   x + width,
			DstY:   y + height,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX: x + width - borderWidth,
			DstY: y + height - borderWidth,
			SrcX: 1,
			SrcY: 1,
		},
	}, borderBoxIndices, brushImage, &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	})
}
