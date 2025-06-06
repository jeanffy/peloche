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
	label        *widget.Label
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewMain(appUIContext *context.AppUIContext) *EditorViewMain {
	x := &EditorViewMain{
		appUIContext: appUIContext,
	}

	x.label = widget.NewLabel("Editor view")
	x.UIContainer = container.NewVBox(x.label)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *EditorViewMain) Activate(fyneWin fyne.Window, args ...interface{}) {
	str := args[0].(*ExplorerViewMainPhotoContainer)
	x.label.SetText(str.photo.Path)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
