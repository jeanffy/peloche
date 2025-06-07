package explorer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerView struct {
	UIContainer fyne.CanvasObject
	main        *ExplorerViewMain
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerView() *ExplorerView {
	x := &ExplorerView{}

	toolbar := NewExplorerViewToolbar()
	leftBar := NewExplorerViewLeftBar()
	x.main = NewExplorerViewMain()

	bottom := container.NewHSplit(leftBar.UIContainer, x.main.UIContainer)
	bottom.Offset = 0.3

	x.UIContainer = container.NewBorder(toolbar.UIContainer, nil, nil, nil, bottom)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerView) Activate() {
	x.main.Activate()
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
