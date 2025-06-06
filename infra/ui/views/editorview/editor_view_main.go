package editorview

import (
	"peloche/domain"
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
	appUIContext *context.UIContext
	label        *widget.Label
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewMain(appUIContext *context.UIContext) *EditorViewMain {
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

func (x *EditorViewMain) Activate(fyneWin fyne.Window, photo *domain.Photo) {
	x.label.SetText(photo.Path)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
