package explorerview

import (
	"peloche/infra/ui/context"
	"peloche/infra/ui/events"
	"peloche/infra/ui/routing"
	"peloche/internal/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMain struct {
	uiContext *context.UIContext
	eventBus  events.EventBus
	router    routing.Router

	UIContainer    fyne.CanvasObject
	progressDialog dialog.Dialog
	photoGrid      *ExplorerViewMainPhotoGrid
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMain() *ExplorerViewMain {
	x := &ExplorerViewMain{
		uiContext: di.GetBasicDI().Resolve(context.UI_CONTEXT_TOKEN).(*context.UIContext),
		eventBus:  di.GetBasicDI().Resolve(events.EVENT_BUS_TOKEN).(events.EventBus),
		router:    di.GetBasicDI().Resolve(routing.ROUTER_TOKEN).(routing.Router),
	}

	toolbar := NewExplorerViewMainToolbar()
	x.photoGrid = NewExplorerViewMainPhotoGrid()

	x.UIContainer = container.NewBorder(nil, toolbar.UIContainer, nil, nil, x.photoGrid.UIContainer)

	x.eventBus.Subscribe(events.EventCurrentFolderChanging, x.onCurrentFolderChanging)
	x.eventBus.Subscribe(events.EventCurrentFolderChanged, x.onCurrentFolderChanged)

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
	currentWin := x.router.GetCurrentWindow()
	x.progressDialog = dialog.NewCustomWithoutButtons("Loading photos...", widget.NewProgressBarInfinite(), currentWin)
	x.progressDialog.Resize(fyne.NewSize(300, 50))
	x.progressDialog.Show()
}

func (x *ExplorerViewMain) onCurrentFolderChanged(event *events.EventCurrentFolderChangedParams) {
	x.progressDialog.Hide()
	x.uiContext.SetSelectedPhotoIndex(0)
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
