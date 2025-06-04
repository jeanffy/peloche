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

type EditorViewMain struct {
	UIContainer  fyne.CanvasObject
	appUIContext *context.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewMain(appUIContext *context.AppUIContext) *EditorViewMain {
	x := &EditorViewMain{
		appUIContext: appUIContext,
	}

	x.UIContainer = container.NewVBox(widget.NewLabel("Editor view"))

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
