package views

import (
	"peloche/infra/ui"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMainToolbar struct {
	UIContainer fyne.CanvasObject

	appUIContext *ui.AppUIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainToolbar(appUIContext *ui.AppUIContext) *ExplorerViewMainToolbar {
	x := &ExplorerViewMainToolbar{
		appUIContext: appUIContext,
	}

	thumbnailSlider := widget.NewSlider(80, 500)
	thumbnailSlider.SetValue(appUIContext.GridSize)

	thumbnailSlider.OnChanged = func(size float64) {
		x.appUIContext.GridSize = size
		x.appUIContext.EventBus.Publish(events.EventThumbnailSizeChanged, &events.EventThumbnailSizeChangedParams{
			Size: size,
		})
	}

	thumbnailSize := fyne.NewSize(150, thumbnailSlider.MinSize().Height)
	x.UIContainer = container.NewHBox(layout.NewSpacer(), container.NewGridWrap(thumbnailSize, thumbnailSlider))

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
