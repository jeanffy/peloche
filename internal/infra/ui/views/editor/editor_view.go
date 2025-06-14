package editor

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui"
	"peloche/pkg/di"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// #region definition

type EditorView struct {
	routerPort ui.RouterPort

	UIContainer fyne.CanvasObject

	main *EditorViewMain
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewEditorView() *EditorView {
	x := &EditorView{
		routerPort: di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort),
	}

	toolbar := NewEditorViewToolbar()
	x.main = NewEditorViewMain()

	x.UIContainer = container.NewBorder(toolbar.UIContainer, nil, nil, nil, x.main.UIContainer)

	return x
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *EditorView) Activate(photo *domain.Photo) {
	x.routerPort.GetCurrentWindow().Canvas().SetOnTypedKey(x.onKeyPress)
	x.main.Activate(photo)
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

func (x *EditorView) onKeyPress(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		x.main.Deactivate()
		runtime.GC()
		x.routerPort.NavigateToExplorerView()
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
