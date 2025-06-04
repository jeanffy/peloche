package views

import (
	"peloche/infra/ui/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type EditorView struct {
	UIContainer  fyne.CanvasObject
	appUIContext *context.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorView(fyneWin fyne.Window, appUIContext *context.AppUIContext) *EditorView {
	x := &EditorView{
		appUIContext: appUIContext,
	}

	toolbar := NewEditorViewToolbar(x.appUIContext)
	main := NewEditorViewMain(x.appUIContext)

	x.UIContainer = container.NewBorder(toolbar.UIContainer, nil, nil, nil, main.UIContainer)

	fyneWin.Canvas().SetOnTypedKey(x.onKeyPress)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *EditorView) onKeyPress(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		x.appUIContext.WinManager.CloseEditorWindow()
	}
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
