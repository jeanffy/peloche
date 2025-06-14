package explorer

import (
	"image/color"
	"peloche/internal/domain"
	"peloche/internal/infra/ui"
	"peloche/internal/infra/ui/assets"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ExplorerViewMainPhotoThumb struct {
	widget.BaseWidget

	context    *ui.Context
	eventsPort ui.EventsPort

	size        float32
	uiContainer *fyne.Container
	photo       *domain.Photo
	image       *canvas.Image
	background  fyne.CanvasObject
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewExplorerViewMainPhotoThumb(photo *domain.Photo, size float32) *ExplorerViewMainPhotoThumb {
	x := &ExplorerViewMainPhotoThumb{
		context:    di.GetBasicDI().Resolve(ui.CONTEXT_TOKEN).(*ui.Context),
		eventsPort: di.GetBasicDI().Resolve(ui.EVENTS_PORT_TOKEN).(ui.EventsPort),
		size:       size,
		photo:      photo,
	}

	x.image = canvas.NewImageFromResource(assets.ResourcePhotoLoadingJpg)
	x.image.FillMode = canvas.ImageFillContain
	x.background = canvas.NewRectangle(color.Transparent)

	x.uiContainer = container.NewCenter()
	x.refreshUIContainer()

	return x
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *ExplorerViewMainPhotoThumb) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(x.uiContainer)
}

func (x *ExplorerViewMainPhotoThumb) LoadImage() {
	x.photo.LoadThumbnailBuffer(x.context.GridSizeMax)
	if x.photo.ThumbnailBuffer == nil {
		x.image = canvas.NewImageFromResource(assets.ResourcePhotoErrorJpg)
	} else {
		x.image = canvas.NewImageFromImage(x.photo.ThumbnailBuffer)
	}
	x.image.FillMode = canvas.ImageFillContain
	fyne.Do(func() {
		x.refreshUIContainer()
	})
}

func (x *ExplorerViewMainPhotoThumb) SetSize(size float32) {
	x.size = size
	x.refreshUIContainer()
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

func (x *ExplorerViewMainPhotoThumb) refreshUIContainer() {
	x.uiContainer.RemoveAll()
	x.uiContainer.Add(container.NewGridWrap(fyne.NewSize(x.size, x.size), x.background))
	x.uiContainer.Add(container.NewGridWrap(fyne.NewSize(x.size-20, x.size-20), x.image))
}

// #endregion
