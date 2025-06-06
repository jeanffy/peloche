package editorview

import (
	"peloche/domain"
	"peloche/infra/ui/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type EditorView struct {
	UIContainer  fyne.CanvasObject
	appUIContext *context.UIContext
	main         *EditorViewMain
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorView(appUIContext *context.UIContext) *EditorView {
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

func (x *EditorView) Activate(fyneWin fyne.Window, photo *domain.Photo) {
	fyneWin.Canvas().SetOnTypedKey(x.onKeyPress)
	x.main.Activate(fyneWin, photo)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *EditorView) onKeyPress(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		x.appUIContext.NavigateToExplorerView()
	}
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
