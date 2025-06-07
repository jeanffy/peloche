package context

import (
	"peloche/internal/infra/ui/events"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

var UI_CONTEXT_TOKEN = "UIContext"

type UIContext struct {
	eventBus events.EventBus

	ThemeVariant       fyne.ThemeVariant
	GridSize           uint
	GridSizeMin        uint
	GridSizeMax        uint
	SelectedPhotoIndex int
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewUIContext(fyneApp fyne.App) *UIContext {
	return &UIContext{
		eventBus:           di.GetBasicDI().Resolve(events.EVENT_BUS_TOKEN).(events.EventBus),
		ThemeVariant:       fyneApp.Settings().ThemeVariant(),
		GridSize:           200,
		GridSizeMin:        80,
		GridSizeMax:        500,
		SelectedPhotoIndex: -1,
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *UIContext) SetGridSize(size uint) {
	x.GridSize = size
	x.eventBus.Publish(events.EventThumbnailSizeChanged, &events.EventThumbnailSizeChangedParams{
		Size: size,
	})
}

func (x *UIContext) SetSelectedPhotoIndex(index int) {
	x.SelectedPhotoIndex = index
	x.eventBus.Publish(events.EventSelectedPhotoChanged, &events.EventSelectedPhotoChangedParams{
		Index: index,
	})
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
