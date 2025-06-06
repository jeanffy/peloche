package ui

import (
	"peloche/domain"
	"peloche/infra/ui/views/editorview"
	"peloche/infra/ui/views/explorerview"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type UIRouter struct {
	fyneWin      fyne.Window
	explorerView *explorerview.ExplorerView
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
		x.explorerView = explorerview.NewExplorerView()
	}
	x.fyneWin.SetContent(x.explorerView.UIContainer)
	x.explorerView.Activate(x.fyneWin)
}

func (x *UIRouter) NavigateToEditorView(photo *domain.Photo) {
	view := editorview.NewEditorView()
	x.fyneWin.SetContent(view.UIContainer)
	view.Activate(x.fyneWin, photo)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
