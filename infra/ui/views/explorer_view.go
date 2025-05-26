package views

import (
	"peloche/infra/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerView struct {
	UIContainer  fyne.CanvasObject
	appUIContext *ui.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerView(appUIContext *ui.AppUIContext) *ExplorerView {
	x := &ExplorerView{
		appUIContext: appUIContext,
	}

	toolbar := NewExplorerViewToolbar(x.appUIContext)
	leftBar := NewExplorerViewLeftBar(x.appUIContext)
	main := NewExplorerViewMain(x.appUIContext)

	bottom := container.NewHSplit(leftBar.UIContainer, main.UIContainer)
	bottom.Offset = 0.3

	x.UIContainer = container.NewBorder(toolbar.UIContainer, nil, nil, nil, bottom)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
