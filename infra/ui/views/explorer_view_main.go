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
	appUIContext   *context.AppUIContext
	progressDialog dialog.Dialog
	photoGrid      *ExplorerViewMainPhotoGrid
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMain(appUIContext *context.AppUIContext) *ExplorerViewMain {
	x := &ExplorerViewMain{
		appUIContext: appUIContext,
	}

	toolbar := NewExplorerViewMainToolbar(x.appUIContext)
	x.photoGrid = NewExplorerViewMainPhotoGrid(x.appUIContext)

	x.UIContainer = container.NewBorder(nil, toolbar.UIContainer, nil, nil, x.photoGrid.UIContainer)

	x.appUIContext.SubscribeToEvent(events.EventCurrentFolderChanging, x.onCurrentFolderChanging)
	x.appUIContext.SubscribeToEvent(events.EventCurrentFolderChanged, x.onCurrentFolderChanged)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerViewMain) Activate(fyneWin fyne.Window) {
	x.photoGrid.Activate(fyneWin)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewMain) onCurrentFolderChanging(event *events.EventCurrentFolderChangingParams) {
	// x.appUIContext.LogInfo(ports.LogPortParams{
	// 	Module: reflect.TypeOf(ExplorerViewMain{}).Name(),
	// 	Msg:    "onCurrentFolderChanging " + event.CurrentFolderPath,
	// })
	currentWin := x.appUIContext.GetCurrentWindow()
	x.progressDialog = dialog.NewCustomWithoutButtons("Loading photos...", widget.NewProgressBarInfinite(), currentWin)
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
