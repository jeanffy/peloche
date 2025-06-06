package explorerview

import (
	"image/color"
	"peloche/domain"
	"peloche/infra/ui/assets"
	"peloche/infra/ui/context"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type ExplorerViewMainPhotoContainer struct {
	UIContainer  fyne.CanvasObject
	appUIContext *context.AppUIContext
	photo        *domain.Photo
	index        int
	image        *canvas.Image
	loaded       bool
	selected     bool
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewExplorerViewMainPhotoContainer(appUIContext *context.AppUIContext, photo *domain.Photo, index int) *ExplorerViewMainPhotoContainer {
	x := &ExplorerViewMainPhotoContainer{
		appUIContext: appUIContext,
		photo:        photo,
		index:        index,
		loaded:       false,
		selected:     false,
	}

	x.UIContainer = container.NewStack()
	x.refreshImageContainer()

	x.appUIContext.SubscribeToEvent(events.EventSelectedPhotoChanged, x.onSelectedPhotoIndexChanged)

	return x
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoContainer) loadBuffer() {
	x.photo.LoadThumbnailBuffer(x.appUIContext.GridSizeMax)
	x.loaded = true
	x.refreshImageContainer()
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoContainer) onSelectedPhotoIndexChanged(event *events.EventSelectedPhotoChangedParams) {
	x.selected = false
	if event.Index == x.index {
		x.selected = true
	}
	x.refreshImageContainer()
}

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

func (x *ExplorerViewMainPhotoContainer) refreshImageContainer() {
	if x.loaded {
		if x.photo.Buffer == nil {
			x.image = canvas.NewImageFromResource(assets.ResourcePhotoErrorJpg)
		} else {
			x.image = canvas.NewImageFromImage(x.photo.Buffer)
		}
	} else {
		x.image = canvas.NewImageFromResource(assets.ResourcePhotoLoadingJpg)
	}

	x.image.FillMode = canvas.ImageFillContain

	var fillColor color.Color
	if x.appUIContext.ThemeVariant == theme.VariantDark {
		fillColor = color.Black
	} else {
		fillColor = color.Gray16{0x888f}
	}

	var selectionRect *canvas.Rectangle = nil
	if x.selected {
		selectionRect = canvas.NewRectangle(color.Transparent)
		selectionRect.StrokeWidth = 3
		themeColorR, themeColorG, themeColorB, themeColorA := theme.Color(theme.ColorNameSelection).RGBA()
		if themeColorA == 0 {
			selectionRect.StrokeColor = color.RGBA{100, 100, 0, 255}
		} else {
			r := (float64(themeColorR) / float64(themeColorA)) * 255.0
			g := (float64(themeColorG) / float64(themeColorA)) * 255.0
			b := (float64(themeColorB) / float64(themeColorA)) * 255.0
			selectionRect.StrokeColor = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}

	fyne.Do(func() {
		x.UIContainer.(*fyne.Container).RemoveAll()
		x.UIContainer.(*fyne.Container).Add(canvas.NewRectangle(fillColor))
		x.UIContainer.(*fyne.Container).Add(x.image)
		if selectionRect != nil {
			x.UIContainer.(*fyne.Container).Add(selectionRect)
		}
		x.UIContainer.Refresh()
	})
}
