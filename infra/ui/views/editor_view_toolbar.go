package views

import (
	"peloche/infra/ui/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type EditorViewToolbar struct {
	UIContainer fyne.CanvasObject

	appUIContext *context.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewToolbar(appUIContext *context.AppUIContext) *EditorViewToolbar {
	instance := &EditorViewToolbar{
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
