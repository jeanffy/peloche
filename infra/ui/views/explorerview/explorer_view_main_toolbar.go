package explorerview

import (
	"peloche/infra/ui/context"

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

	appUIContext *context.UIContext
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainToolbar(appUIContext *context.UIContext) *ExplorerViewMainToolbar {
	x := &ExplorerViewMainToolbar{
		appUIContext: appUIContext,
	}

	thumbnailSlider := widget.NewSlider(float64(appUIContext.GridSizeMin), float64(appUIContext.GridSizeMax))
	thumbnailSlider.SetValue(float64(appUIContext.GridSize))

	thumbnailSlider.OnChanged = func(size float64) {
		x.appUIContext.SetGridSize(uint(size))
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
