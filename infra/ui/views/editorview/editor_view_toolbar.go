package editorview

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

type EditorViewToolbar struct {
	dialogs dialogs.Dialogs

	UIContainer fyne.CanvasObject
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewToolbar() *EditorViewToolbar {
	x := &EditorViewToolbar{
		dialogs: utils.GetNaiveDI().Resolve(dialogs.DIALOGS_TOKEN).(dialogs.Dialogs),
	}

	button1 := widget.NewButton("Some toolbar button", func() {
		x.dialogs.MessageDialog("Some toolbar button clicked")
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
