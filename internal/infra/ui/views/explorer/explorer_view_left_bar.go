package explorer

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui/context"
	"peloche/internal/infra/ui/dialogs"
	"peloche/internal/infra/ui/events"
	"peloche/internal/infra/ui/routing"
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
	uiContext *context.UIContext
	eventBus  events.EventBus
	router    routing.Router
	dialogs   dialogs.Dialogs
	appData   *domain.AppData

	UIContainer fyne.CanvasObject

	openFolderButton *widget.Button
	tree             *widget.Tree
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewLeftBar() *ExplorerViewLeftBar {
	x := &ExplorerViewLeftBar{
		uiContext: di.GetBasicDI().Resolve(context.UI_CONTEXT_TOKEN).(*context.UIContext),
		eventBus:  di.GetBasicDI().Resolve(events.EVENT_BUS_TOKEN).(events.EventBus),
		router:    di.GetBasicDI().Resolve(routing.ROUTER_TOKEN).(routing.Router),
		dialogs:   di.GetBasicDI().Resolve(dialogs.DIALOGS_TOKEN).(dialogs.Dialogs),
		appData:   di.GetBasicDI().Resolve(domain.APP_DATA_TOKEN).(*domain.AppData),
	}

	x.openFolderButton = widget.NewButton(lang.L("views.explorer.openFolder"), x.onOpenFolderClicked)

	folderTree := NewExplorerViewLeftBarFolderTree()
	x.tree = folderTree.UIContainer.(*widget.Tree)
	x.tree.Hide()

	x.UIContainer = container.NewBorder(container.NewHBox(x.openFolderButton), nil, nil, nil, x.tree)

	x.eventBus.Subscribe(events.EventRootFolderChanged, x.onRootFolderChanged)

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
			x.dialogs.ErrorDialog(err)
		}
		return
	}

	x.appData.SetRootFolder(&directory)

	x.eventBus.Publish(events.EventRootFolderChanged, &events.EventRootFolderChangedParams{
		RootFolderPath: directory,
		FolderTree:     x.appData.FolderTree,
	})
}

func (x *ExplorerViewLeftBar) onRootFolderChanged(event *events.EventRootFolderChangedParams) {
	fyne.Do(func() {
		x.tree.Show()
		x.UIContainer.Refresh()
	})
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
