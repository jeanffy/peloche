package views

import (
	"peloche/infra/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewToolbar struct {
	UIContainer fyne.CanvasObject

	appUIContext *ui.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewToolbar(appUIContext *ui.AppUIContext) *ExplorerViewToolbar {
	instance := &ExplorerViewToolbar{
		appUIContext: appUIContext,
	}

	button1 := widget.NewButton("Some toolbar button", func() {
		dialog.Message("%s", "Some toolbar button clicked").Info()
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
