package ui

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui/views/editor"
	"peloche/internal/infra/ui/views/explorer"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type UIRouter struct {
	fyneWin      fyne.Window
	explorerView *explorer.ExplorerView
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewUIRouter(fyneWin fyne.Window) *UIRouter {
	return &UIRouter{
		fyneWin: fyneWin,
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *UIRouter) GetCurrentWindow() fyne.Window {
	return x.fyneWin
}

func (x *UIRouter) NavigateToExplorerView() {
	if x.explorerView == nil {
		x.explorerView = explorer.NewExplorerView()
	}
	x.fyneWin.SetContent(x.explorerView.UIContainer)
	x.explorerView.Activate(x.fyneWin)
}

func (x *UIRouter) NavigateToEditorView(photo *domain.Photo) {
	view := editor.NewEditorView()
	x.fyneWin.SetContent(view.UIContainer)
	view.Activate(x.fyneWin, photo)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
