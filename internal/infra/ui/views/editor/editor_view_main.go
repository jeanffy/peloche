package editor

import (
	"peloche/internal/domain"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type EditorViewMain struct {
	UIContainer fyne.CanvasObject

	label *widget.Label
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorViewMain() *EditorViewMain {
	x := &EditorViewMain{}

	x.label = widget.NewLabel("Editor view")
	x.UIContainer = container.NewVBox(x.label)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *EditorViewMain) Activate(photo *domain.Photo) {
	x.label.SetText(photo.Path)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
