package explorer

import (
	"peloche/internal/domain"
	"peloche/internal/infra/ui"
	"peloche/internal/infra/ui/fyneex"
	"peloche/pkg/di"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

/*
  +--------------------------+
  | gridScrollContainer      |
  | +----------------------+ |
  | | gridLayout           | |
  | | +------------------+ | |
  | | | gridContainer    | | |
  | | +------------------+ | |
  | +----------------------+ |
  +--------------------------+
*/

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type ExplorerViewMainPhotoGrid struct {
	context    *ui.Context
	eventsPort ui.EventsPort
	appData    *domain.AppData
	routerPort ui.RouterPort

	UIContainer fyne.CanvasObject

	photoContainers []*ExplorerViewMainPhotoContainer
	scrollContainer *container.Scroll
	layout          *fyneex.GridWrapLayout
	grid            *fyne.Container
	currentRow      int
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainPhotoGrid() *ExplorerViewMainPhotoGrid {
	x := &ExplorerViewMainPhotoGrid{
		context:         di.GetBasicDI().Resolve(ui.CONTEXT_TOKEN).(*ui.Context),
		eventsPort:      di.GetBasicDI().Resolve(ui.EVENTS_PORT_TOKEN).(ui.EventsPort),
		appData:         di.GetBasicDI().Resolve(domain.APP_DATA_TOKEN).(*domain.AppData),
		routerPort:      di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort),
		photoContainers: []*ExplorerViewMainPhotoContainer{},
		currentRow:      0,
	}

	// we should be using a GridWrap widget here but unfortunately, the fyne GridWrap widget
	// "swallows" key press events with its TypedKey event
	// so as a workaround, we use a GridWrap layout
	x.scrollContainer = container.NewScroll(container.NewGridWrap(fyne.NewSize(0, 0)))
	x.createLayout()

	x.UIContainer = x.scrollContainer

	x.eventsPort.Subscribe(ui.EventCurrentFolderChanged, x.onCurrentFolderChanged)
	x.eventsPort.Subscribe(ui.EventThumbnailSizeChanged, x.onThumbnailSizeChanged)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoGrid) Activate() {
	x.routerPort.GetCurrentWindow().Canvas().SetOnTypedKey(x.onKeyPress)
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoGrid) onCurrentFolderChanged(event *ui.EventCurrentFolderChangedParams) {
	x.photoContainers = make([]*ExplorerViewMainPhotoContainer, len(x.appData.PhotoList.Photos))
	for i, photo := range x.appData.PhotoList.Photos {
		x.photoContainers[i] = NewExplorerViewMainPhotoContainer(photo, i)
	}

	x.buildGridWithPhotos()

	go func() {
		for _, photo := range x.photoContainers {
			photo.loadBuffer()
			fyne.Do(func() {
				if photo.index < len(x.grid.Objects) {
					x.grid.Objects[photo.index] = fyneex.NewClickableThing(photo, photo.UIContainer, x.onPhotoSelected)
				}
			})
		}
	}()
}

func (x *ExplorerViewMainPhotoGrid) onPhotoSelected(photo *ExplorerViewMainPhotoContainer) {
	currentIndex := x.context.SelectedPhotoIndex
	if photo.index != currentIndex {
		x.context.SetSelectedPhotoIndex(photo.index)
	}
}

func (x *ExplorerViewMainPhotoGrid) onThumbnailSizeChanged(event *ui.EventThumbnailSizeChangedParams) {
	x.createLayout()
	x.buildGridWithPhotos()
}

func (x *ExplorerViewMainPhotoGrid) onKeyPress(key *fyne.KeyEvent) {
	if slices.Contains([]fyne.KeyName{fyne.KeyUp, fyne.KeyDown, fyne.KeyLeft, fyne.KeyRight}, key.Name) {
		x.onArrowKeyPressed(key.Name)
	} else if key.Name == fyne.KeySpace {
		x.onSpaceBarPressed()
	}
}

func (x *ExplorerViewMainPhotoGrid) onArrowKeyPressed(keyName fyne.KeyName) {
	indexMax := len(x.photoContainers) - 1
	currentIndex := x.context.SelectedPhotoIndex
	var nextIndex = currentIndex

	if keyName == fyne.KeyLeft {
		nextIndex = nextIndex - 1
	} else if keyName == fyne.KeyRight {
		nextIndex = nextIndex + 1
	} else if keyName == fyne.KeyUp {
		nextIndex -= x.layout.ColCount
		if nextIndex < 0 {
			nextIndex = currentIndex
		}
	} else if keyName == fyne.KeyDown {
		nextIndex += x.layout.ColCount
		if nextIndex > indexMax {
			nextIndex = currentIndex
		}
	}

	if nextIndex < 0 {
		nextIndex = 0
	} else if nextIndex > indexMax {
		nextIndex = indexMax
	}

	if nextIndex != currentIndex {
		x.context.SetSelectedPhotoIndex(nextIndex)
		// // TODO: scroll to the selected photo
		// x.currentRow = int(x.layout.ColCount / nextIndex)
		// fmt.Println(x.currentRow)
		// x.scrollContainer.ScrollToOffset(fyne.NewPos(0, float32(x.currentRow*int(x.appUIContext.GridSize))))
	}
}

func (x *ExplorerViewMainPhotoGrid) onSpaceBarPressed() {
	if x.context.SelectedPhotoIndex != -1 {
		x.editPhoto(x.photoContainers[x.context.SelectedPhotoIndex])
	}
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoGrid) createLayout() {
	size := float32(x.context.GridSize)
	x.layout = fyneex.NewMyGridWrapLayout(fyne.NewSize(size, size)).(*fyneex.GridWrapLayout)
	x.grid = container.New(x.layout)
	x.scrollContainer.Content = x.grid
}

func (x *ExplorerViewMainPhotoGrid) buildGridWithPhotos() {
	x.grid.RemoveAll()
	for _, photo := range x.photoContainers {
		x.grid.Add(fyneex.NewClickableThing(photo, photo.UIContainer, x.onPhotoSelected))
	}
	x.scrollContainer.Refresh()
}

func (x *ExplorerViewMainPhotoGrid) editPhoto(photoContainer *ExplorerViewMainPhotoContainer) {
	x.routerPort.NavigateToEditorView(photoContainer.photo)
}
