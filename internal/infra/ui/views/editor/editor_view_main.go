package editor

import (
	"peloche/internal/domain"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// ---------------------------------------------------------------------------
// #region definition

type EditorViewMain struct {
	UIContainer fyne.CanvasObject
	photo       *domain.Photo
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewEditorViewMain() *EditorViewMain {
	x := &EditorViewMain{}

	x.UIContainer = container.NewStack()

	return x
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *EditorViewMain) Activate(photo *domain.Photo) {
	x.photo = photo

	x.photo.LoadBuffer()

	image := canvas.NewImageFromImage(x.photo.Buffer)
	image.FillMode = canvas.ImageFillContain

	x.UIContainer.(*fyne.Container).RemoveAll()
	x.UIContainer.(*fyne.Container).Add(image)
	x.UIContainer.Refresh()
}

func (x *EditorViewMain) Deactivate() {
	x.photo.FreeBuffer()
	x.photo = nil
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
