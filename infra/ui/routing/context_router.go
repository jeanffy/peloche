package routing

import (
	"peloche/domain"

	"fyne.io/fyne/v2"
)

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
