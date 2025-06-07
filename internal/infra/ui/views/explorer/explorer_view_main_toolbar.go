package explorer

import (
	"peloche/internal/infra/ui"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMainToolbar struct {
	context *ui.Context

	UIContainer fyne.CanvasObject
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainToolbar() *ExplorerViewMainToolbar {
	x := &ExplorerViewMainToolbar{
		context: di.GetBasicDI().Resolve(ui.CONTEXT_TOKEN).(*ui.Context),
	}

	thumbnailSlider := widget.NewSlider(float64(x.context.GridSizeMin), float64(x.context.GridSizeMax))
	thumbnailSlider.SetValue(float64(x.context.GridSize))

	thumbnailSlider.OnChanged = func(size float64) {
		x.context.SetGridSize(uint(size))
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
