package editor

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui/routing"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type EditorView struct {
	router routing.Router

	UIContainer fyne.CanvasObject

	main *EditorViewMain
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewEditorView() *EditorView {
	x := &EditorView{
		router: di.GetBasicDI().Resolve(routing.ROUTER_TOKEN).(routing.Router),
	}

	toolbar := NewEditorViewToolbar()
	x.main = NewEditorViewMain()

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
		x.router.NavigateToExplorerView()
	}
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
