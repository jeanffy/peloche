package explorerview

import (
	"peloche/infra/ui/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewToolbar struct {
	UIContainer fyne.CanvasObject

	appUIContext *context.UIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewToolbar(appUIContext *context.UIContext) *ExplorerViewToolbar {
	instance := &ExplorerViewToolbar{
		appUIContext: appUIContext,
	}

	button1 := widget.NewButton("Some toolbar button", func() {
		appUIContext.ShowMessageBox("Some toolbar button clicked")
	})
	instance.UIContainer = container.NewHBox(button1)

	return instance
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
