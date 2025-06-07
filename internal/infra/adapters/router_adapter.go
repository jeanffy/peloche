package adapters

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui/views/editor"
	"peloche/internal/infra/ui/views/explorer"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type RouterAdapter struct {
	fyneWin      fyne.Window
	explorerView *explorer.ExplorerView
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewRouterAdapter(fyneWin fyne.Window) *RouterAdapter {
	return &RouterAdapter{
		fyneWin: fyneWin,
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
