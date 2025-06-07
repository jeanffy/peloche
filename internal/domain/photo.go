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

func (x *Photo) LoadThumbnailBuffer(thumbnailSize uint) {
	if x.ThumbnailBuffer == nil {
		img := x.getDecodedPhoto(x.Path, x.Ext)
		x.ThumbnailBuffer = resize.Thumbnail(thumbnailSize, thumbnailSize, img, resize.NearestNeighbor)
		img = nil
		runtime.GC()
	}
}

func (x *Photo) LoadBuffer() {
	if x.Buffer == nil {
		x.Buffer = x.getDecodedPhoto(x.Path, x.Ext)
	}
}

func (x *Photo) FreeBuffer() {
	x.Buffer = nil
}

func (x *Photo) getDecodedPhoto(filePath string, ext string) image.Image {
	reader, err := os.Open(filePath)
	if err != nil {
		x.log.Error(LogPortErrorParams{
			Module: reflect.TypeOf(PhotoList{}).Name(),
			Error:  err,
			Msg:    filePath,
		})
		return nil
	}

	defer reader.Close()

	var imgDecoded image.Image = nil

	if strings.ToLower(ext) == ".heic" {
		img, err := goheif.Decode(reader)
		if err != nil {
			x.log.Error(LogPortErrorParams{
				Module: reflect.TypeOf(PhotoList{}).Name(),
				Error:  err,
				Msg:    filePath,
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
				Msg:    filePath,
			})
			return nil
		}
		imgDecoded = img
	}

	return imgDecoded
}
