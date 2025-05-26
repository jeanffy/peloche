package views

import (
	"fmt"
	"peloche/domain"
	"peloche/domain/ports"
	"peloche/infra/ui"
	"peloche/infra/ui/events"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

/*
for a folder tree like:

root/
+--foo/
|  +--bar/
+--baz/
   +--dummy

ids:
"":         {"root"}
"root":     {"root/foo", "root/baz"}
"root/foo": {"root/foo/bar"}
"root/baz": {"root/baz/dummy"}

values:
"root":           "root"
"root/foo":       "foo"
"root/baz":       "baz"
"root/foo/bar":   "bar"
"root/baz/dummy": "dummy"
*/

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewLeftBarFolderTree struct {
	UIContainer fyne.CanvasObject

	appUIContext *ui.AppUIContext

	ids    map[string][]string
	values map[string]string

	treeBinding binding.ExternalStringTree
	tree        *widget.Tree
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewLeftBarFolderTree(appUIContext *ui.AppUIContext) *ExplorerViewLeftBarFolderTree {
	x := &ExplorerViewLeftBarFolderTree{
		appUIContext: appUIContext,
		ids:          map[string][]string{},
		values:       map[string]string{},
	}

	x.treeBinding = binding.BindStringTree(&x.ids, &x.values)

	x.tree = widget.NewTreeWithData(
		x.treeBinding,
		// CreateNode
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		// UpdateNode
		func(data binding.DataItem, branch bool, o fyne.CanvasObject) {
			d, err := data.(binding.String).Get()
			if err != nil {
				return
			}
			text := d
			// if branch {
			// 	text += " (branch)"
			// }
			o.(*widget.Label).SetText(text)
		},
	)

	x.tree.OnSelected = func(id string) {
		folder := x.appUIContext.AppData.FolderTree.Find(id)
		if folder == nil {
			x.appUIContext.AppData.Log.Error(ports.LogPortErrorParams{
				Module: reflect.TypeOf(ExplorerViewLeftBarFolderTree{}).Name(),
				Error:  fmt.Errorf("no folder found for id '%s'", id),
			})
			return
		}
		x.onTreeItemClicked(folder)
	}

	x.UIContainer = x.tree

	x.appUIContext.EventBus.Subscribe(events.EventRootFolderChanged, x.onRootFolderChanged)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewLeftBarFolderTree) onRootFolderChanged(event *events.EventRootFolderChangedParams) {
	ids := map[string][]string{}
	values := map[string]string{}
	createTreeBindingIds(x.appUIContext.AppData.FolderTree, &ids, true)
	createTreeBindingValues(x.appUIContext.AppData.FolderTree, &values)
	x.treeBinding.Set(ids, values)
	x.tree.OpenBranch(ids[""][0])
	x.tree.Refresh()
}

func (x *ExplorerViewLeftBarFolderTree) onTreeItemClicked(folder *domain.FolderTree) {
	folderPath := folder.Path

	x.appUIContext.EventBus.Publish(events.EventCurrentFolderChanging, &events.EventCurrentFolderChangingParams{
		CurrentFolderPath: folderPath,
	})

	// setting the current folder in a go routine to keep the UI reactive
	// and to let the previous event be published to subscribers
	go func() {
		x.appUIContext.AppData.SetCurrentFolder(&folderPath)

		x.appUIContext.EventBus.Publish(events.EventCurrentFolderChanged, &events.EventCurrentFolderChangedParams{
			CurrentFolderPath: folderPath,
			PhotoList:         x.appUIContext.AppData.PhotoList,
		})
	}()
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

func createTreeBindingIds(folderTree *domain.FolderTree, ids *map[string][]string, isRoot bool) {
	subFolderIds := make([]string, 0, len(folderTree.SubFolders))
	if isRoot {
		(*ids)[""] = []string{folderTree.Path}
	}
	for _, subFolder := range folderTree.SubFolders {
		subFolderIds = append(subFolderIds, subFolder.Path)
		createTreeBindingIds(subFolder, ids, false)
	}
	(*ids)[folderTree.Path] = subFolderIds
}

func createTreeBindingValues(folderTree *domain.FolderTree, values *map[string]string) {
	(*values)[folderTree.Path] = folderTree.Name
	for _, subFolder := range folderTree.SubFolders {
		(*values)[subFolder.Path] = subFolder.Name
		createTreeBindingValues(subFolder, values)
	}
}
