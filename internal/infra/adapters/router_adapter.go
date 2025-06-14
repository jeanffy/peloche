package adapters

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui"
	"peloche/internal/infra/ui/views/editor"
	"peloche/internal/infra/ui/views/explorer"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// #region definition

var _ ui.RouterPort = (*RouterAdapter)(nil)

type RouterAdapter struct {
	fyneWin      fyne.Window
	explorerView *explorer.ExplorerView
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewRouterAdapter(fyneWin fyne.Window) *RouterAdapter {
	return &RouterAdapter{
		fyneWin: fyneWin,
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *RouterAdapter) GetCurrentWindow() fyne.Window {
	return x.fyneWin
}

func (x *RouterAdapter) NavigateToExplorerView() {
	if x.explorerView == nil {
		x.explorerView = explorer.NewExplorerView()
	}
	x.fyneWin.SetContent(x.explorerView.UIContainer)
	x.explorerView.Activate()
}

func (x *RouterAdapter) NavigateToEditorView(photo *domain.Photo) {
	view := editor.NewEditorView()
	x.fyneWin.SetContent(view.UIContainer)
	view.Activate(photo)
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
