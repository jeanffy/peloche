package explorer

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"

	xdialog "github.com/sqweek/dialog"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewLeftBar struct {
	uiContext   *ui.Context
	eventsPort  ui.EventsPort
	routerPort  ui.RouterPort
	dialogsPort ui.DialogsPort
	appData     *domain.AppData

	UIContainer fyne.CanvasObject

	openFolderButton *widget.Button
	tree             *widget.Tree
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewLeftBar() *ExplorerViewLeftBar {
	x := &ExplorerViewLeftBar{
		uiContext:   di.GetBasicDI().Resolve(ui.CONTEXT_TOKEN).(*ui.Context),
		eventsPort:  di.GetBasicDI().Resolve(ui.EVENTS_PORT_TOKEN).(ui.EventsPort),
		routerPort:  di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort),
		dialogsPort: di.GetBasicDI().Resolve(ui.DIALOGS_PORT_TOKEN).(ui.DialogsPort),
		appData:     di.GetBasicDI().Resolve(domain.APP_DATA_TOKEN).(*domain.AppData),
	}

	x.openFolderButton = widget.NewButton(lang.L("views.explorer.openFolder"), x.onOpenFolderClicked)

	folderTree := NewExplorerViewLeftBarFolderTree()
	x.tree = folderTree.UIContainer.(*widget.Tree)
	x.tree.Hide()

	x.UIContainer = container.NewBorder(container.NewHBox(x.openFolderButton), nil, nil, nil, x.tree)

	x.eventsPort.Subscribe(ui.EventRootFolderChanged, x.onRootFolderChanged)

	return x

}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewLeftBar) onOpenFolderClicked() {
	directory, err := xdialog.Directory().Title(lang.L("views.explorer.openFolder")).Browse()
	if err != nil {
		if err != xdialog.ErrCancelled {
			x.dialogsPort.ErrorDialog(err)
		}
		return
	}

	x.appData.SetRootFolder(&directory)

	x.eventsPort.Publish(ui.EventRootFolderChanged, &ui.EventRootFolderChangedParams{
		RootFolderPath: directory,
		FolderTree:     x.appData.FolderTree,
	})
}

func (x *ExplorerViewLeftBar) onRootFolderChanged(event *ui.EventRootFolderChangedParams) {
	fyne.Do(func() {
		x.tree.Show()
		x.UIContainer.Refresh()
	})
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
