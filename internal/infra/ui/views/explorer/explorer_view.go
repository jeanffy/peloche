package explorer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// #region definition

type ExplorerView struct {
	UIContainer fyne.CanvasObject
	main        *ExplorerViewMain
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

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

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *ExplorerView) Activate() {
	x.main.Activate()
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
