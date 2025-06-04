package views

import (
	"peloche/infra/ui/context"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMain struct {
	UIContainer    fyne.CanvasObject
	fyneWin        fyne.Window
	appUIContext   *context.AppUIContext
	progressDialog dialog.Dialog
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMain(fyneWin fyne.Window, appUIContext *context.AppUIContext) *ExplorerViewMain {
	x := &ExplorerViewMain{
		appUIContext: appUIContext,
		fyneWin:      fyneWin,
	}

	toolbar := NewExplorerViewMainToolbar(x.appUIContext)
	photoGrid := NewExplorerViewMainPhotoGrid(fyneWin, x.appUIContext)

	x.UIContainer = container.NewBorder(nil, toolbar.UIContainer, nil, nil, photoGrid.UIContainer)

	x.appUIContext.EventBus.Subscribe(events.EventCurrentFolderChanging, x.onCurrentFolderChanging)
	x.appUIContext.EventBus.Subscribe(events.EventCurrentFolderChanged, x.onCurrentFolderChanged)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewMain) onCurrentFolderChanging(event *events.EventCurrentFolderChangingParams) {
	// x.appUIContext.LogInfo(ports.LogPortParams{
	// 	Module: reflect.TypeOf(ExplorerViewMain{}).Name(),
	// 	Msg:    "onCurrentFolderChanging " + event.CurrentFolderPath,
	// })
	x.progressDialog = dialog.NewCustomWithoutButtons("Loading photos...", widget.NewProgressBarInfinite(), x.fyneWin)
	x.progressDialog.Resize(fyne.NewSize(300, 50))
	x.progressDialog.Show()
}

func (x *ExplorerViewMain) onCurrentFolderChanged(event *events.EventCurrentFolderChangedParams) {
	x.progressDialog.Hide()
	x.appUIContext.SetSelectedPhotoIndex(0)
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
