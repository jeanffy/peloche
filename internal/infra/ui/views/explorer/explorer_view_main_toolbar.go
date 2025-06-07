package explorer

import (
	"peloche/internal/infra/ui/context"
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
	uiContext *context.UIContext

	UIContainer fyne.CanvasObject
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainToolbar() *ExplorerViewMainToolbar {
	x := &ExplorerViewMainToolbar{
		uiContext: di.GetBasicDI().Resolve(context.UI_CONTEXT_TOKEN).(*context.UIContext),
	}

	thumbnailSlider := widget.NewSlider(float64(x.uiContext.GridSizeMin), float64(x.uiContext.GridSizeMax))
	thumbnailSlider.SetValue(float64(x.uiContext.GridSize))

	thumbnailSlider.OnChanged = func(size float64) {
		x.uiContext.SetGridSize(uint(size))
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
