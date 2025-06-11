package explorer

import (
	"fmt"
	"peloche/internal/domain"
	"peloche/internal/infra/ui"
	"peloche/pkg/di"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMainPhotoGrid struct {
	context    *ui.Context
	eventsPort ui.EventsPort
	appData    *domain.AppData
	routerPort ui.RouterPort

	UIContainer *fyne.Container

	gridSize      float32
	data          []any
	dataBinding   binding.ExternalUntypedList
	grid          *widget.GridWrap
	photoThumbs   map[string]*ExplorerViewMainPhotoThumb
	selectedIndex int
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainPhotoGrid() *ExplorerViewMainPhotoGrid {
	x := &ExplorerViewMainPhotoGrid{
		context:       di.GetBasicDI().Resolve(ui.CONTEXT_TOKEN).(*ui.Context),
		eventsPort:    di.GetBasicDI().Resolve(ui.EVENTS_PORT_TOKEN).(ui.EventsPort),
		appData:       di.GetBasicDI().Resolve(domain.APP_DATA_TOKEN).(*domain.AppData),
		routerPort:    di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort),
		selectedIndex: -1,
	}

	if x.appData.PhotoList == nil {
		return x
	}

	x.gridSize = float32(x.context.GridSize)

	domainPhotos := x.appData.PhotoList.Photos

	x.photoThumbs = map[string]*ExplorerViewMainPhotoThumb{}
	for _, photo := range domainPhotos {
		x.photoThumbs[photo.Name] = NewExplorerViewMainPhotoThumb(photo, x.gridSize)
	}

	// taken from https://stackoverflow.com/a/73029665
	var unpackArray = func(s any) []any {
		v := reflect.ValueOf(s)
		r := make([]any, v.Len())
		for i := 0; i < v.Len(); i++ {
			r[i] = v.Index(i).Interface()
		}
		return r
	}

	x.data = unpackArray(domainPhotos)
	x.dataBinding = binding.BindUntypedList(&x.data)
	x.grid = x.createGrid()

	x.UIContainer = container.NewStack(x.grid)

	x.eventsPort.Subscribe(ui.EventThumbnailSizeChanged, x.onThumbnailSizeChanged)

	go func() {
		for _, photoThumb := range x.photoThumbs {
			photoThumb.LoadImage()
		}
	}()

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoGrid) Activate() {
	ctrlE := &desktop.CustomShortcut{
		Modifier: fyne.KeyModifierShortcutDefault,
		KeyName:  fyne.KeyE,
	}
	x.routerPort.GetCurrentWindow().Canvas().AddShortcut(ctrlE, func(shortcut fyne.Shortcut) {
		fmt.Println("edit")
	})
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoGrid) onThumbnailSizeChanged(event *ui.EventThumbnailSizeChangedParams) {
	x.gridSize = float32(event.Size)
	oldSelectedIndex := x.selectedIndex
	x.grid = x.createGrid()
	if oldSelectedIndex != -1 {
		x.grid.Select(oldSelectedIndex)
		x.grid.ScrollTo(oldSelectedIndex)
	}
	x.UIContainer.RemoveAll()
	x.UIContainer.Add(x.grid)
}

// func (x *ExplorerViewMainPhotoGrid) onSpaceBarPressed() {
// 	if x.context.SelectedPhotoIndex != -1 {
// 		x.editPhoto(x.photoContainers[x.context.SelectedPhotoIndex])
// 	}
// }

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoGrid) createGrid() *widget.GridWrap {
	grid := widget.NewGridWrapWithData(
		x.dataBinding,
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(x.gridSize, x.gridSize))
		},
		func(data binding.DataItem, o fyne.CanvasObject) {
			d, err := data.(binding.Untyped).Get()
			if err != nil {
				return
			}
			photo := d.(*domain.Photo)
			photoThumbnail := x.photoThumbs[photo.Name]
			photoThumbnail.SetSize(x.gridSize)
			wrapper := o.(*fyne.Container)
			wrapper.RemoveAll()
			wrapper.Add(photoThumbnail)
		},
	)
	grid.OnSelected = func(id widget.GridWrapItemID) {
		x.selectedIndex = id
	}
	grid.OnUnselected = func(id widget.GridWrapItemID) {
		x.selectedIndex = -1
	}
	return grid
}

func (x *ExplorerViewMainPhotoGrid) editPhoto(photoContainer *ExplorerViewMainPhotoThumb) {
	x.routerPort.NavigateToEditorView(photoContainer.photo)
}
