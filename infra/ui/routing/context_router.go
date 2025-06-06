package routing

import "fyne.io/fyne/v2"

type RouteName int8

const (
	RouteExplorer RouteName = iota
	RouteEditor
)

type Router interface {
	GetCurrentWindow() fyne.Window
	NavigateTo(route RouteName, args ...interface{})
}
