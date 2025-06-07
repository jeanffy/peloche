package ui

import (
	"peloche/internal/domain"

	"fyne.io/fyne/v2"
)

var ROUTER_PORT_TOKEN = "RouterPort"

type RouteName int8

const (
	RouteExplorer RouteName = iota
	RouteEditor
)

type RouterPort interface {
	GetCurrentWindow() fyne.Window
	NavigateToExplorerView()
	NavigateToEditorView(photo *domain.Photo)
}
