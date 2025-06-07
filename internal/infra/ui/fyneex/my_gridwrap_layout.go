//
// /-------------------------------------------------------------------------\
// |                                                                         |
// | taken and modified from 2.6.0 version layout/gridwraplayout.go          |
// | purpose: get the ColCount and the RowCount of the GridWrap layout       |
// | modifications to the original file are tagged with the "JPC:" marker    |
// |                                                                         |
// \-------------------------------------------------------------------------/
//

package fyneex

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*GridWrapLayout)(nil)

// JPC: uppercase G to export struct
type GridWrapLayout struct {
	CellSize fyne.Size
	ColCount int // JPC: uppercase C to export property
	RowCount int // JPC: uppercase R to export property
}

// NewGridWrapLayout returns a new GridWrapLayout instance
func NewMyGridWrapLayout(size fyne.Size) fyne.Layout {
	return &GridWrapLayout{size, 1, 1}
}

// Layout is called to pack all child objects into a specified size.
// For a GridWrapLayout this will attempt to lay all the child objects in a row
// and wrap to a new row if the size is not large enough.
func (g *GridWrapLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	padding := theme.Padding()
	g.ColCount = 1
	g.RowCount = 0

	if size.Width > g.CellSize.Width {
		g.ColCount = int(math.Floor(float64(size.Width+padding) / float64(g.CellSize.Width+padding)))
	}

	i, x, y := 0, float32(0), float32(0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if i%g.ColCount == 0 {
			g.RowCount++
		}

		child.Move(fyne.NewPos(x, y))
		child.Resize(g.CellSize)

		if (i+1)%g.ColCount == 0 {
			x = 0
			y += g.CellSize.Height + padding
		} else {
			x += g.CellSize.Width + padding
		}
		i++
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For a GridWrapLayout this is simply the specified cellsize as a single column
// layout has no padding. The returned size does not take into account the number
// of columns as this layout re-flows dynamically.
func (g *GridWrapLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	rows := g.RowCount
	if rows < 1 {
		rows = 1
	}
	return fyne.NewSize(g.CellSize.Width,
		(g.CellSize.Height*float32(rows))+(float32(rows-1)*theme.Padding()))
}
