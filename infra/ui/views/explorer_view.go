package views

import (
	"peloche/infra/ui/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerView struct {
	UIContainer  fyne.CanvasObject
	appUIContext *context.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerView(fyneWin fyne.Window, appUIContext *context.AppUIContext) *ExplorerView {
	x := &ExplorerView{
		appUIContext: appUIContext,
	}

	toolbar := NewExplorerViewToolbar(x.appUIContext)
	leftBar := NewExplorerViewLeftBar(x.appUIContext)
	main := NewExplorerViewMain(fyneWin, x.appUIContext)

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
