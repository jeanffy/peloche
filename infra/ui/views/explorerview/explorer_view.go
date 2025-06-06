package explorerview

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
	appUIContext *context.UIContext
	main         *ExplorerViewMain
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerView(appUIContext *context.UIContext) *ExplorerView {
	x := &ExplorerView{
		appUIContext: appUIContext,
	}

	toolbar := NewExplorerViewToolbar(x.appUIContext)
	leftBar := NewExplorerViewLeftBar(x.appUIContext)
	x.main = NewExplorerViewMain(x.appUIContext)

	bottom := container.NewHSplit(leftBar.UIContainer, x.main.UIContainer)
	bottom.Offset = 0.3

	x.UIContainer = container.NewBorder(toolbar.UIContainer, nil, nil, nil, bottom)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerView) Activate(fyneWin fyne.Window, args ...interface{}) {
	x.main.Activate(fyneWin)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
