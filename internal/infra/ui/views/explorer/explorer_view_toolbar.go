package explorer

import (
	"peloche/internal/infra/ui"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewToolbar struct {
	UIContainer fyne.CanvasObject
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewToolbar() *ExplorerViewToolbar {
	x := &ExplorerViewToolbar{}

	dialogsPort := di.GetBasicDI().Resolve(ui.DIALOGS_PORT_TOKEN).(ui.DialogsPort)

	button1 := widget.NewButton("Some toolbar button", func() {
		dialogsPort.MessageDialog("Some toolbar button clicked")
	})
	x.UIContainer = container.NewHBox(button1)

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
