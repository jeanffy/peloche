package explorerview

import (
	"peloche/infra/ui/dialogs"
	"peloche/utils"

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
	instance := &ExplorerViewToolbar{}

	dialogs := utils.GetNaiveDI().Resolve(dialogs.DIALOGS_TOKEN).(dialogs.Dialogs)

	button1 := widget.NewButton("Some toolbar button", func() {
		dialogs.MessageDialog("Some toolbar button clicked")
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
