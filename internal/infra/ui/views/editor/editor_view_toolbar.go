package editor

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

type EditorViewToolbar struct {
	dialogsPort ui.DialogsPort

	UIContainer fyne.CanvasObject
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewToolbar() *EditorViewToolbar {
	x := &EditorViewToolbar{
		dialogsPort: di.GetBasicDI().Resolve(ui.DIALOGS_PORT_TOKEN).(ui.DialogsPort),
	}

	button1 := widget.NewButton("Some toolbar button", func() {
		x.dialogsPort.MessageDialog("Some toolbar button clicked")
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
