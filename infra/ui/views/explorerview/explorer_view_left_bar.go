package explorerview

import (
	"peloche/domain"
	"peloche/infra/ui/context"
	"peloche/infra/ui/dialogs"
	"peloche/infra/ui/events"
	"peloche/infra/ui/routing"
	"peloche/utils"

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
		uiContext: utils.GetNaiveDI().Resolve(context.UI_CONTEXT_TOKEN).(*context.UIContext),
		eventBus:  utils.GetNaiveDI().Resolve(events.EVENT_BUS_TOKEN).(events.EventBus),
		router:    utils.GetNaiveDI().Resolve(routing.ROUTER_TOKEN).(routing.Router),
		dialogs:   utils.GetNaiveDI().Resolve(dialogs.DIALOGS_TOKEN).(dialogs.Dialogs),
		appData:   utils.GetNaiveDI().Resolve(domain.APP_DATA_TOKEN).(*domain.AppData),
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
