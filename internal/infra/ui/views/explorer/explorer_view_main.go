package explorer

import (
	"peloche/internal/infra/ui"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMain struct {
	context    *ui.Context
	eventsPort ui.EventsPort
	routerPort ui.RouterPort

	UIContainer        fyne.CanvasObject
	progressDialog     dialog.Dialog
	photoGrid          *ExplorerViewMainPhotoGrid
	photoGridContainer *fyne.Container
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMain() *ExplorerViewMain {
	x := &ExplorerViewMain{
		context:    di.GetBasicDI().Resolve(ui.CONTEXT_TOKEN).(*ui.Context),
		eventsPort: di.GetBasicDI().Resolve(ui.EVENTS_PORT_TOKEN).(ui.EventsPort),
		routerPort: di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort),
	}

	toolbar := NewExplorerViewMainToolbar()

	x.photoGrid = NewExplorerViewMainPhotoGrid()
	x.photoGridContainer = container.NewStack()
	x.UIContainer = container.NewBorder(nil, toolbar.UIContainer, nil, nil, x.photoGridContainer)

	x.eventsPort.Subscribe(ui.EventCurrentFolderChanging, x.onCurrentFolderChanging)
	x.eventsPort.Subscribe(ui.EventCurrentFolderChanged, x.onCurrentFolderChanged)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerViewMain) Activate() {
	x.photoGrid.Activate()
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewMain) onCurrentFolderChanging(event *ui.EventCurrentFolderChangingParams) {
	// x.appUIContext.LogInfo(domain.LogPortParams{
	// 	Module: reflect.TypeOf(ExplorerViewMain{}).Name(),
	// 	Msg:    "onCurrentFolderChanging " + event.CurrentFolderPath,
	// })
	currentWin := x.routerPort.GetCurrentWindow()
	x.progressDialog = dialog.NewCustomWithoutButtons("Loading photos...", widget.NewProgressBarInfinite(), currentWin)
	x.progressDialog.Resize(fyne.NewSize(300, 50))
	x.progressDialog.Show()
}

func (x *ExplorerViewMain) onCurrentFolderChanged(event *ui.EventCurrentFolderChangedParams) {
	x.progressDialog.Hide()
	x.context.SetSelectedPhotoIndex(0)
	x.photoGrid = NewExplorerViewMainPhotoGrid()
	x.photoGridContainer.RemoveAll()
	x.photoGridContainer.Add(x.photoGrid.UIContainer)
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
