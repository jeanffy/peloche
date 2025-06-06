package ui

import (
	"peloche/infra/ui/context"
	"peloche/infra/ui/views"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type AppUIRouter struct {
	fyneWin      fyne.Window
	appUIContext *context.AppUIContext
	explorerView *views.ExplorerView
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewAppUIRouter(fyneWin fyne.Window) *AppUIRouter {
	return &AppUIRouter{
		fyneWin: fyneWin,
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *AppUIRouter) SetAppUIContext(appUIContext *context.AppUIContext) {
	x.appUIContext = appUIContext
}

func (x *AppUIRouter) GetCurrentWindow() fyne.Window {
	return x.fyneWin
}

func (x *AppUIRouter) NavigateTo(route context.Route) {
	if route == context.RouteExplorer {
		x.goToExplorerView()
	} else if route == context.RouteEditor {
		x.goToEditorView()
	}
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

func (x *AppUIRouter) goToExplorerView() {
	if x.explorerView == nil {
		x.explorerView = views.NewExplorerView(x.appUIContext)
	}
	x.fyneWin.SetContent(x.explorerView.UIContainer)
	x.explorerView.Activate(x.fyneWin)
}

func (x *AppUIRouter) goToEditorView() {
	view := views.NewEditorView(x.appUIContext)
	x.fyneWin.SetContent(view.UIContainer)
	view.Activate(x.fyneWin)
}
