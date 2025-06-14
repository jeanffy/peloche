package ui

import (
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// #region definition

var CONTEXT_TOKEN = "Context"

type Context struct {
	eventsPort EventsPort

	ThemeVariant       fyne.ThemeVariant
	GridSize           uint
	GridSizeMin        uint
	GridSizeMax        uint
	SelectedPhotoIndex int
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewContext(fyneApp fyne.App) *Context {
	return &Context{
		eventsPort:         di.GetBasicDI().Resolve(EVENTS_PORT_TOKEN).(EventsPort),
		ThemeVariant:       fyneApp.Settings().ThemeVariant(),
		GridSize:           200,
		GridSizeMin:        80,
		GridSizeMax:        500,
		SelectedPhotoIndex: -1,
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *Context) SetGridSize(size uint) {
	x.GridSize = size
	x.eventsPort.Publish(EventThumbnailSizeChanged, &EventThumbnailSizeChangedParams{
		Size: size,
	})
}

func (x *Context) SetSelectedPhotoIndex(index int) {
	x.SelectedPhotoIndex = index
	x.eventsPort.Publish(EventSelectedPhotoChanged, &EventSelectedPhotoChangedParams{
		Index: index,
	})
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
