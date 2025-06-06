package context

import "fyne.io/fyne/v2"

type Route int8

const (
	RouteExplorer Route = iota
	RouteEditor
)

type ContextRouter interface {
	GetCurrentWindow() fyne.Window
	NavigateTo(route Route)
}
