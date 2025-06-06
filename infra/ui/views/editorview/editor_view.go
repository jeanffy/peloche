package editorview

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
	main         *EditorViewMain
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorView(appUIContext *context.AppUIContext) *EditorView {
	x := &EditorView{
		appUIContext: appUIContext,
	}

	toolbar := NewEditorViewToolbar(x.appUIContext)
	x.main = NewEditorViewMain(x.appUIContext)

	x.UIContainer = container.NewBorder(toolbar.UIContainer, nil, nil, nil, x.main.UIContainer)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *EditorView) Activate(fyneWin fyne.Window, args ...interface{}) {
	fyneWin.Canvas().SetOnTypedKey(x.onKeyPress)
	x.main.Activate(fyneWin, args...)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *EditorView) onKeyPress(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		x.appUIContext.NavigateTo(context.RouteExplorer)
	}
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
