package ui

import (
	"peloche/infra/ui/context"
	"peloche/infra/ui/routing"
	"peloche/infra/ui/views/editorview"
	"peloche/infra/ui/views/explorerview"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type AppUIRouter struct {
	fyneWin      fyne.Window
	appUIContext *context.UIContext
	explorerView *explorerview.ExplorerView
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

func (x *AppUIRouter) SetAppUIContext(appUIContext *context.UIContext) {
	x.appUIContext = appUIContext
}

func (x *AppUIRouter) GetCurrentWindow() fyne.Window {
	return x.fyneWin
}

func (x *AppUIRouter) NavigateTo(route routing.RouteName, args ...interface{}) {
	if route == routing.RouteExplorer {
		x.goToExplorerView(args...)
	} else if route == routing.RouteEditor {
		x.goToEditorView(args...)
	}
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

func (x *AppUIRouter) goToExplorerView(args ...interface{}) {
	if x.explorerView == nil {
		x.explorerView = explorerview.NewExplorerView(x.appUIContext)
	}
	x.fyneWin.SetContent(x.explorerView.UIContainer)
	x.explorerView.Activate(x.fyneWin, args...)
}

func (x *AppUIRouter) goToEditorView(args ...interface{}) {
	view := editorview.NewEditorView(x.appUIContext)
	x.fyneWin.SetContent(view.UIContainer)
	view.Activate(x.fyneWin, args...)
}
