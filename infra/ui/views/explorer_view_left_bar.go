package views

import (
	"peloche/infra/ui"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"

	xdialog "github.com/sqweek/dialog"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewLeftBar struct {
	UIContainer fyne.CanvasObject

	appUIContext *ui.AppUIContext

	openFolderButton *widget.Button
	tree             *widget.Tree
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewLeftBar(appUIContext *ui.AppUIContext) *ExplorerViewLeftBar {
	x := &ExplorerViewLeftBar{
		appUIContext: appUIContext,
	}

	x.openFolderButton = widget.NewButton(lang.L("views.explorer.openFolder"), x.onOpenFolderClicked)

	folderTree := NewExplorerViewLeftBarFolderTree(x.appUIContext)
	x.tree = folderTree.UIContainer.(*widget.Tree)
	x.tree.Hide()

	x.UIContainer = container.NewBorder(container.NewHBox(x.openFolderButton), nil, nil, nil, x.tree)

	x.appUIContext.EventBus.Subscribe(events.EventRootFolderChanged, x.onRootFolderChanged)

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
			dialog.NewError(err, x.appUIContext.FyneWin).Show()
		}
		return
	}

	x.appUIContext.AppData.SetRootFolder(&directory)

	x.appUIContext.EventBus.Publish(events.EventRootFolderChanged, &events.EventRootFolderChangedParams{
		RootFolderPath: directory,
		FolderTree:     x.appUIContext.AppData.FolderTree,
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
