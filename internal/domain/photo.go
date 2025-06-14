package domain

import (
	"image"
	"os"
	"peloche/pkg/di"
	"reflect"
	"runtime"
	"strings"

	"github.com/adrium/goheif"
	"github.com/nfnt/resize"
)

// ---------------------------------------------------------------------------
// #region definition

type Photo struct {
	log LogPort

	Name            string
	Ext             string
	Path            string
	Width           int
	Height          int
	ThumbnailBuffer image.Image
	Buffer          image.Image
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewPhoto(name string, ext string, filePath string) *Photo {
	return &Photo{
		log:             di.GetBasicDI().Resolve(LOG_PORT_TOKEN).(LogPort),
		Name:            name,
		Ext:             ext,
		Path:            filePath,
		Width:           0,
		Height:          0,
		ThumbnailBuffer: nil,
		Buffer:          nil,
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *Photo) LoadThumbnailBuffer(thumbnailSize uint) {
	if x.ThumbnailBuffer == nil {
		img := x.getImageBuffer()
		x.ThumbnailBuffer = resize.Thumbnail(thumbnailSize, thumbnailSize, img, resize.NearestNeighbor)
		img = nil
		runtime.GC()
	}
}

func (x *Photo) LoadBuffer() {
	if x.Buffer == nil {
		x.Buffer = x.getImageBuffer()
	}
}

func (x *Photo) FreeBuffer() {
	x.Buffer = nil
}

// #endregion

// ---------------------------------------------------------------------------
// #region private

func (x *Photo) getImageBuffer() image.Image {
	reader, err := os.Open(x.Path)
	if err != nil {
		x.log.Error(LogPortErrorParams{
			Module: reflect.TypeOf(PhotoList{}).Name(),
			Error:  err,
			Msg:    x.Path,
		})
		return nil
	}

	defer reader.Close()

	var imgDecoded image.Image = nil

	if strings.ToLower(x.Ext) == ".heic" {
		img, err := goheif.Decode(reader)
		if err != nil {
			x.log.Error(LogPortErrorParams{
				Module: reflect.TypeOf(PhotoList{}).Name(),
				Error:  err,
				Msg:    x.Path,
			})
			return nil
		}
		imgDecoded = img
	} else {
		img, _, err := image.Decode(reader)
		if err != nil {
			x.log.Error(LogPortErrorParams{
				Module: reflect.TypeOf(PhotoList{}).Name(),
				Error:  err,
				Msg:    x.Path,
			})
			return nil
		}
		imgDecoded = img
	}

	return imgDecoded
}

// #endregion
