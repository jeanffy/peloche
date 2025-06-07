package fyneex

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ClickableThing[T comparable] struct {
	widget.BaseWidget

	thing  T
	uiItem fyne.CanvasObject

	OnTapped func(T)
}

func NewClickableThing[T comparable](thing T, uiItem fyne.CanvasObject, tapped func(T)) *ClickableThing[T] {
	x := &ClickableThing[T]{
		thing:    thing,
		uiItem:   uiItem,
		OnTapped: tapped,
	}
	x.ExtendBaseWidget(x)
	return x
}

func (x *ClickableThing[T]) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(x.uiItem)
}

func (x *ClickableThing[T]) Tapped(*fyne.PointEvent) {
	if x.OnTapped != nil {
		x.OnTapped(x.thing)
	}
}
