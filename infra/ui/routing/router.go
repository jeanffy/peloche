package routing

import (
	"peloche/domain"

	"fyne.io/fyne/v2"
)

var ROUTER_TOKEN = "Router"

type RouteName int8

const (
	RouteExplorer RouteName = iota
	RouteEditor
)

type Router interface {
	GetCurrentWindow() fyne.Window
	NavigateToExplorerView()
	NavigateToEditorView(photo *domain.Photo)
}
